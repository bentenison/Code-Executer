package mid

import (
	"github.com/bentenison/microservice/foundation/otel"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

// Otel starts the otel tracing and stores the trace id in the context.
func Otel(tracer trace.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := otel.InjectTracing(c.Request.Context(), tracer)
		c.Request = c.Request.WithContext(ctx)
		otel.AddTraceToRequest(c.Request.Context(), c.Request)
		// log.Println("tracerID from otel", c.Request.Context().Value(otel.TraceIDKey))
		// log.Println("tracer himself", c.Request.Context().Value(otel.TracerKey))
		// log.Println(c.Request.Context().Value(traceKey))
		c.Next()

	}
}
