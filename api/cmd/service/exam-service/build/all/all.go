package all

import (
	examapi "github.com/bentenison/microservice/api/domain/exam-api"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/exambus"
	"github.com/bentenison/microservice/business/domain/exambus/stores/examdb"
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
	exambus := exambus.NewBusiness(cfg.Log, delegate, examdb.NewStore(cfg.DB, cfg.Log))
	// // Construct the business domain packages we need here so we are using the
	// // sames instances for the different set of domain apis.
	examapi.Routes(app, examapi.Config{
		Log:     cfg.Log,
		ExamBus: exambus,
		// Tracer:    cfg.Tracer,
		AppConfig: cfg.AppConfig,
	})

}
