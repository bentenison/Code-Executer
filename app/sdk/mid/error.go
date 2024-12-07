package mid

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/bentenison/microservice/foundation/otel"
	"github.com/gin-gonic/gin"
)

// ErrorResponse defines a consistent error response format
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	TraceID    string `json:"trace_id"`
}

// ErrorMiddleware is a centralized error handler for the Gin framework
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// Handle panic gracefully
			if err := recover(); err != nil {
				handleError(c, fmt.Errorf("%v", err), http.StatusInternalServerError)
			}
		}()

		// Process the request
		c.Next()

		// Check if any errors were raised during the request
		if len(c.Errors) > 0 {
			lastError := c.Errors.Last()
			statusCode := http.StatusInternalServerError
			if lastError.Meta != nil {
				if code, ok := lastError.Meta.(int); ok {
					statusCode = code // Use custom status code if available
				}
			}
			handleError(c, lastError.Err, statusCode)
		}
	}
}

// handleError handles an individual error, attaches it to the trace, and sends a consistent response
func handleError(c *gin.Context, err error, statusCode int) {
	// Get the OpenTelemetry trace span from the context
	log.Println("error middleware called")
	otel.RecordError(c, statusCode, err)
	// Log stack trace for debugging purposes
	stackTrace := string(debug.Stack())

	// Print the error and stack trace (can be sent to logs or error tracking system)
	fmt.Printf("ERROR: %v\nSTACK TRACE:\n%s\n", err, stackTrace)

	// Get the trace ID from the span
	traceID := otel.GetTraceID(c.Request.Context())

	// Create a standard error response
	response := ErrorResponse{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode),
		Error:      err.Error(),
		TraceID:    traceID,
	}

	// Return the error response to the client
	c.JSON(statusCode, response)
}
