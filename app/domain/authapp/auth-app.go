package authapp

import (
	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/bentenison/microservice/foundation/logger"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	authbus *authbus.Business
	logger  *logger.CustomLogger
	tracer  trace.Tracer
}

func NewApp(log *logger.CustomLogger, authbus *authbus.Business, tp *tp.TracerProvider) *App {
	return &App{
		authbus: authbus,
		logger:  log,
		tracer:  tp.Tracer("AUTH"),
	}
}
