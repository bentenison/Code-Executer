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
	// cfg.Log.Error("started serving ")
	app.Handle("GET", "/broker/getTemplate", api.getQuestionTemplate)
	app.Handle("GET", "/broker/authenticate", api.getQuestionTemplate)
}
