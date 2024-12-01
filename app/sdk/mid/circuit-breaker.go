package mid

import (
	"net/http"
	"time"

	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

const (
	failureThreshold = 3                // Number of failures before opening the circuit
	successThreshold = 2                // Number of successes to close the circuit
	resetTimeout     = 30 * time.Second // Timeout before transitioning from open to half-open
)

type CircuitBreaker struct {
	circuitBreaker *gobreaker.CircuitBreaker
	log            *logger.CustomLogger
}

func NewCircuitBreaker(servicename string, log *logger.CustomLogger) *CircuitBreaker {
	// Define circuit breaker settings
	settings := gobreaker.Settings{
		Name:        servicename,  // Name for the circuit breaker
		Timeout:     resetTimeout, // Timeout for resetting the circuit breaker
		MaxRequests: 5,            // Max requests to send to a service in half-open state
		Interval:    time.Second,  // Interval for metrics and health checks
		// Failure threshold and success threshold to open/close the circuit
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Circuit breaker opens after 'failureThreshold' failures and closes after 'successThreshold' successes
			if counts.ConsecutiveFailures > failureThreshold {
				return true
			}
			if counts.ConsecutiveSuccesses > successThreshold {
				return false
			}
			return false
		},
	}

	// Create and return the CircuitBreaker
	return &CircuitBreaker{circuitBreaker: gobreaker.NewCircuitBreaker(settings), log: log}
}

// CircuitBreakerMiddleware is the Gin middleware that wraps the CircuitBreaker
func CircuitBreakerMiddleware(cb *CircuitBreaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the downstream function call in the circuit breaker logic
		res, err := cb.circuitBreaker.Execute(func() (interface{}, error) {
			c.Next()
			if len(c.Errors) > 0 {
				return nil, c.Errors.Last()
			}
			return nil, nil
		})
		_ = res
		// Check for errors and handle circuit breaker state
		if err != nil {
			if err == gobreaker.ErrOpenState {
				cb.log.Errorc(c.Request.Context(), "error: Circuit breaker is open state", map[string]interface{}{})
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"error": "Service is temporarily unavailable",
				})
				c.Abort()
			} else if err == gobreaker.ErrTooManyRequests {
				cb.log.Errorc(c.Request.Context(), "error: Circuit breaker is in half-open state", map[string]interface{}{})
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": "Too many requests, please try again later",
				})
				c.Abort()
			} else {
				cb.log.Errorc(c.Request.Context(), "error: request failed", map[string]interface{}{
					"error": err.Error(),
				})
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				c.Abort()
			}
		}
	}
}
