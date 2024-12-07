package adminapp

import (
	"github.com/bentenison/microservice/business/domain/adminbus"
	"github.com/bentenison/microservice/foundation/logger"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	adminbus *adminbus.Business
	logger   *logger.CustomLogger
	tracer   trace.Tracer
}

func NewApp(logger *logger.CustomLogger, bus *adminbus.Business) *App {
	return &App{logger: logger, adminbus: bus}
}
