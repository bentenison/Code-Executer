package authapi

import (
	"github.com/bentenison/microservice/app/domain/authapp"
	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Log *logger.CustomLogger
	// authapp
	AuthBus *authbus.Business
	Tracer  *trace.TracerProvider
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(cfg.Log, authapp.NewApp(cfg.Log, cfg.AuthBus, cfg.Tracer))
	app.Handle("GET", "/auth/check", api.checkHealthHandler)
}
