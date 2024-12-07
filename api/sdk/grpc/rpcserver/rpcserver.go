package rpcserver

import (
	"context"
	"net"

	"github.com/bentenison/microservice/api/sdk/grpc/mid"
	"github.com/bentenison/microservice/foundation/logger"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateServer(grpcPort string, log *logger.CustomLogger, tp *trace.TracerProvider) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Errorc(context.TODO(), "failed to listen:", map[string]interface{}{
			"error": err.Error(),
		})
		panic(err)
	}
	reqMid := mid.NewRequestIDMiddleware(log)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(reqMid.UnaryRequestIDInterceptor(), mid.UnaryOtelInterceptor(tp.Tracer(""))),
		grpc.ChainStreamInterceptor(reqMid.StreamRequestIDInterceptor(), mid.StreamOtelInterceptor(tp.Tracer(""))))
	return grpcServer, lis
}

func CreateClient(log *logger.CustomLogger, port string) *grpc.ClientConn {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorc(context.TODO(), "failed to connect:", map[string]interface{}{
			"error": err.Error(),
		})
		panic(err)
		// log.Fatalf("Failed to connect: %v", err)
	}
	return conn
}
