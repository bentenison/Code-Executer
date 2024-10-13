package brokerapi

import (
	brokerapp "github.com/bentenison/microservice/app/domain/broker-app"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Log       *logger.CustomLogger
	BrokerBus *brokerbus.Business
	Tracer    *trace.TracerProvider
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(brokerapp.NewApp(cfg.Log, cfg.BrokerBus, cfg.Tracer), cfg.Log)
	app.Use(mid.TraceIdMiddleware())
	// cfg.Log.Errorc("started serving ")
	app.Handle("POST", "/broker/submission", api.newSubmissionHandler)
	app.Handle("POST", "/broker/batchProcess", api.newSubmissionHandler)
	app.Handle("GET", "/broker/authenticate", api.newSubmissionHandler)
	app.Handle("GET", "/broker/getlanguages", api.newSubmissionHandler)
	app.Handle("GET", "/broker/getallowedlanguages", api.newSubmissionHandler)
	app.Handle("GET", "/broker/conf", api.newSubmissionHandler)
	app.Handle("GET", "/broker/updateconf", api.newSubmissionHandler)
}
