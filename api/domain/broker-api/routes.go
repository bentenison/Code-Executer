package brokerapi

import (
	authpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto"
	execpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto"
	"github.com/bentenison/microservice/api/sdk/grpc/rpcserver"
	brokerapp "github.com/bentenison/microservice/app/domain/broker-app"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Log            *logger.CustomLogger
	BrokerBus      *brokerbus.Business
	Tracer         *trace.TracerProvider
	AppConfig      *conf.Config
	AuthClient     authpb.AuthServiceClient
	ExecutorClient execpb.ExecutorServiceClient
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(brokerapp.NewApp(cfg.Log, cfg.BrokerBus, cfg.Tracer), cfg.Log)
	app.Use(mid.TraceIdMiddleware())

	authcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.AuthGRPCPort)
	authcli := authpb.NewAuthServiceClient(authcliConn)
	api.authcli = authcli
	executorcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.GRPCPort)
	execcli := execpb.NewExecutorServiceClient(executorcliConn)
	api.execcli = execcli
	// cfg.Log.Errorc("started serving ")
	app.Handle("POST", "/broker/submission", api.newSubmissionHandler)
	app.Handle("POST", "/broker/batchProcess", api.newSubmissionHandler)
	app.Handle("GET", "/broker/authenticate", api.newSubmissionHandler)
	app.Handle("GET", "/broker/getlanguages", api.newSubmissionHandler)
	app.Handle("GET", "/broker/getallowedlanguages", api.newSubmissionHandler)
	app.Handle("GET", "/broker/conf", api.newSubmissionHandler)
	app.Handle("GET", "/broker/updateconf", api.newSubmissionHandler)
}
