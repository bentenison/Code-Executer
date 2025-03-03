package mid

import (
	"context"

	"github.com/bentenison/microservice/foundation/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func UnaryOtelInterceptor(tracer trace.Tracer) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Retrieve request ID from incoming metadata or create a new on
		// fmt.Println("I am getting the requestId here", ctx.Value("tracectx"))
		ctx = otel.InjectTracing(ctx, tracer)
		// Add the request ID to the context

		// Proceed with the request
		return handler(ctx, req)
	}
}
func StreamOtelInterceptor(tracer trace.Tracer) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {

		// ctx := context.Background()
		// Add the request ID to the context
		wrappedStream := &serverStreamWithContext{
			ServerStream: ss,
			ctx:          otel.InjectTracing(ss.Context(), tracer),
		}

		// Proceed with the stream
		return handler(srv, wrappedStream)
	}
}
