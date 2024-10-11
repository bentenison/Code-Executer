package all

import (
	brokerapi "github.com/bentenison/microservice/api/domain/broker-api"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/business/domain/brokerbus/stores/brokerdb"
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
	brokerbus := brokerbus.NewBusiness(cfg.Log, cfg.DB, delegate, brokerdb.NewStore(cfg.DB, cfg.Log))
	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	brokerapi.Routes(app, brokerapi.Config{
		Log:       cfg.Log,
		BrokerBus: brokerbus,
		Tracer:    cfg.Tracer,
	})

}
