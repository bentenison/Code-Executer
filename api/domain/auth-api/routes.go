package authapi

import (
	"github.com/bentenison/microservice/app/domain/authapp"
	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/trace"
)

type Config struct {
	Log *logger.CustomLogger
	// authapp
	AuthBus *authbus.Business
	Tracer  *trace.TracerProvider
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(cfg.log, authapp.NewApp(cfg.log, cfg.AuthBus))
	app.Handle("GET", "/auth/check", api.checkHealthHandler)
}
