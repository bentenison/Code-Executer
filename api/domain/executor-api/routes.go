package executorapi

import (
	"github.com/bentenison/microservice/app/domain/executorapp"
	"github.com/bentenison/microservice/business/domain/executorbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Log         *logger.CustomLogger
	Tracer      *trace.TracerProvider
	ExecutorBus *executorbus.Business
	// DockerClient *client.Client
}

func Routes(app *web.App, conf Config) {
	api := newAPI(executorapp.NewApp(conf.ExecutorBus, conf.Log, conf.Tracer), conf.Log)

	app.Handle("POST", "executor/handlesubmission", api.handleSubmission)
}
