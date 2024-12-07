package otel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	trc "go.opentelemetry.io/otel/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.13.0"
)

type Config struct {
	Host        string
	Probability float64
	ServiceName string
}

func NewTracer(conf Config) (*trace.TracerProvider, error) {
	return initTracer(conf)
}
func initTracer(conf Config) (*trace.TracerProvider, error) {
	// Create a stdout exporter to export the trace data to the console.

	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Host)))
	if err != nil {
		return nil, fmt.Errorf("failed to create Jaeger exporter: %v", err)
	}

	// tp := tracesdk.NewTracerProvider(
	// 	trace.WithBatcher(exporter),
	// 	trace.WithSampler(trace.TraceIDRatioBased(0.1)),
	// )

	// otel.SetTracerProvider(tp)
	// return tp, nil
	// exporter, err := jaeger.Wi(jaeger.WithAgentEndpoint("localhost:5775")) // Change to Docker Jaeger container's IP if necessary
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create Jaeger exporter: %v", err)
	// }

	// Set up a batch span processor to push traces to Jaeger
	// bsp := trace.NewBatchSpanProcessor(jaegerExporter)
	// exporter, err := otlptrace.New(
	// 	context.Background(),
	// 	otlptracegrpc.NewClient(
	// 		otlptracegrpc.WithEndpointURL("http://localhost:4317"),
	// 		otlptracegrpc.WithInsecure(), // No TLS
	// 	),
	// )
	// if err != nil {
	// 	return nil, fmt.Errorf("creating new exporter: %w", err)
	// }

	// Create a TracerProvider with a batcher and resource attributes
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(conf.Probability))),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(conf.ServiceName),
			),
		),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func ShutDownTracer(tp *trace.TracerProvider) { _ = tp.Shutdown(context.Background()) }
func InjectTracing(ctx context.Context, tracer trc.Tracer) context.Context {
	ctx = setTracer(ctx, tracer)

	traceID := trc.SpanFromContext(ctx).SpanContext().TraceID().String()
	if traceID == "00000000000000000000000000000000" {
		traceID = uuid.NewString()
	}
	ctx = setTraceID(ctx, traceID)

	return ctx
}

// AddSpan adds an otel span to the existing trace.
func AddSpan(ctx context.Context, spanName string, keyValues ...attribute.KeyValue) (context.Context, trc.Span) {
	v, ok := ctx.Value(TracerKey).(trc.Tracer)
	if !ok || v == nil {
		return ctx, trc.SpanFromContext(ctx)
	}

	ctx, span := v.Start(ctx, spanName)
	for _, kv := range keyValues {
		span.SetAttributes(kv)
	}

	return ctx, span
}

func RecordError(c *gin.Context, statusCode int, err error) {

	// Attach the error details to the trace
	// v, ok := c.Request.Context().Value(TracerKey).(trc.Tracer)
	// if !ok || v == nil {
	// 	return
	// }

	_, span := AddSpan(c.Request.Context(), c.Request.URL.Path)
	span.RecordError(err)
	span.SetAttributes(
		attribute.String("error.message", err.Error()),
		attribute.String("http.path", c.Request.URL.Path),
		attribute.String("http.method", c.Request.Method),
		attribute.Int("http.status_code", statusCode),
	)
	span.End()

}

// AddTraceToRequest adds the current trace id to the request so it
// can be delivered to the service being called.
func AddTraceToRequest(ctx context.Context, r *http.Request) {
	hc := propagation.HeaderCarrier(r.Header)
	otel.GetTextMapPropagator().Inject(ctx, hc)
}
