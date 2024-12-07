package examapp

import (
	"github.com/bentenison/microservice/business/domain/exambus"
	"github.com/bentenison/microservice/foundation/logger"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	brokerbus *exambus.Business
	logger    *logger.CustomLogger
	tracer    trace.Tracer
}

func NewApp(logger *logger.CustomLogger, bus *exambus.Business) *App {
	return &App{logger: logger, brokerbus: bus}
}
