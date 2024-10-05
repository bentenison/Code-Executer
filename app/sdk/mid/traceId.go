package mid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationID := c.GetHeader("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.NewString()
		}
		c.Set(traceKey, correlationID)
		c.Next()
	}
}
