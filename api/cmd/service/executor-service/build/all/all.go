package all

import (
	executorapi "github.com/bentenison/microservice/api/domain/executor-api"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/executorbus"
	"github.com/bentenison/microservice/business/domain/executorbus/stores/executordb"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/docker/docker/client"
)

func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	delegate := delegate.New(cfg.Log)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	execbus := executorbus.NewBusiness(cfg.Log, delegate, executordb.NewStore(cfg.Log, cfg.DB), cli)
	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.

	executorapi.Routes(app, executorapi.Config{
		Log:         cfg.Log,
		ExecutorBus: execbus,
		Tracer:      cfg.Tracer,
		AppConfig:   cfg.AppConfig,
		// DockerClient: cli,
	})

}
