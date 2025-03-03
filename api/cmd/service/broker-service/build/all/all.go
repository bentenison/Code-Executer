package all

import (
	brokerapi "github.com/bentenison/microservice/api/domain/broker-api"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/adminclient/proto/admClient"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto/execClient"
	"github.com/bentenison/microservice/api/sdk/grpc/rpcserver"
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
	// crete grp  clients for injection
	// authcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.AuthGRPCPort)
	// authcli := authpb.NewAuthServiceClient(authcliConn)
	executorcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.GRPCPort)
	adminclient := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.AdminGRPCPort)
	execcli := execClient.NewExecutorServiceClient(executorcliConn)
	admincli := admClient.NewAdminServiceClient(adminclient)
	delegate := delegate.New(cfg.Log)
	brokerbus := brokerbus.NewBusiness(cfg.Log, delegate, brokerdb.NewStore(cfg.DB, cfg.Log), cfg.RabbitProducer, admincli)
	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	brokerapi.Routes(app, brokerapi.Config{
		Log:            cfg.Log,
		BrokerBus:      brokerbus,
		Tracer:         cfg.Tracer,
		AppConfig:      cfg.AppConfig,
		ExecutorClient: execcli,
		AdminClient:    admincli,
	})

}
