package all

import (
	creatorapi "github.com/bentenison/microservice/api/domain/creator-api"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/creatorbus"
	"github.com/bentenison/microservice/business/domain/creatorbus/stores/creatordb"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/web"
)

func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	delegate := delegate.New(cfg.Log)
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	panic(err)
	// }
	creatorbus := creatorbus.NewBusiness(cfg.Log, delegate, creatordb.NewStore(cfg.DB, cfg.Log))
	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	creatorapi.Routes(app, creatorapi.Config{
		Log:        cfg.Log,
		CreatorBus: creatorbus,
		// Tracer:     cfg.Tracer,
		AppConfig: cfg.AppConfig,
		// DockerClient: cli,
	})

}
