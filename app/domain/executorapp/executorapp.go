package executorapp

import (
	"context"

	pb "github.com/bentenison/microservice/api/domain/executor-api/grpc/proto"
	"github.com/bentenison/microservice/business/domain/executorbus"
	"github.com/bentenison/microservice/foundation/logger"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	executorbus *executorbus.Business
	log         *logger.CustomLogger
	tracer      trace.Tracer
	// dockerClient *client.Client
}

func NewApp(executorbus *executorbus.Business, log *logger.CustomLogger, tracer *tp.TracerProvider) *App {
	return &App{
		executorbus: executorbus,
		log:         log,
		tracer:      tracer.Tracer("EXECUTOR"),
		// dockerClient: cli,
	}
}

func (a *App) HandleExecution(ctx context.Context, path, language, qid, uid string) (*pb.ExecutionResponse, error) {
	return a.executorbus.ExecuteCode(ctx, path, language, uid, qid)

}
