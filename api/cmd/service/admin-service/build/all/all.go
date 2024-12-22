package all

import (
	adminapi "github.com/bentenison/microservice/api/domain/admin-api"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/adminbus"
	"github.com/bentenison/microservice/business/domain/adminbus/stores/admindb"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	delegate := delegate.New(cfg.Log)
	adminBus := adminbus.NewBusiness(cfg.Log, delegate, admindb.NewStore(cfg.DB, cfg.Log), cfg.RabbitConsumer)
	// // Construct the business domain packages we need here so we are using the
	// // sames instances for the different set of domain apis.
	adminapi.Routes(app, adminapi.Config{
		Log:            cfg.Log,
		AdminBus:       adminBus,
		Tracer:         cfg.Tracer,
		AppConfig:      cfg.AppConfig,
		KafkaConsumer:  cfg.KafkaConsumer,
		RabbitConsumer: cfg.RabbitConsumer,
	})

}
