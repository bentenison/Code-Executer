package brokerapp

import (
	"context"

	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	brokerbus *brokerbus.Business
	logger    *logger.CustomLogger
	tracer    trace.Tracer
}

func NewApp(logger *logger.CustomLogger, bus *brokerbus.Business, tp *tp.TracerProvider) *App {
	return &App{logger: logger, brokerbus: bus, tracer: tp.Tracer("BROKER")}
}

func (a *App) HandleSubmisson(ctx context.Context, submission Submission) (brokerbus.Question, error) {
	return a.brokerbus.HandleSubmissonService(ctx, brokerbus.Submission(submission))
}
