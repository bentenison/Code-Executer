package mux

import (
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/sdk/trace"
)

type RouteAdder interface {
	Add(app *web.App, cfg Config)
}
type DataSource struct {
	SQL *sqlx.DB
	MGO *mongo.Database
}
type Config struct {
	Build string
	Log   *logger.CustomLogger
	// Auth       *auth.Auth
	// AuthClient *authclient.Client
	DB     DataSource
	Tracer *trace.TracerProvider
}

func WebAPI(cfg Config, routeAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Log, cfg.Build)
	routeAdder.Add(app, cfg)
	return app
}
