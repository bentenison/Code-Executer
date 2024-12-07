package adminapi

import (
	"net/http"
	"time"

	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/adminbus"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type Config struct {
	Log       *logger.CustomLogger
	AdminBus  *adminbus.Business
	Tracer    *trace.TracerProvider
	AppConfig *conf.Config
}

func Routes(app *web.App, cfg Config) {
	// authcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.AuthGRPCPort)
	// authcli := authpb.NewAuthServiceClient(authcliConn)
	// executorcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.GRPCPort)
	// execcli := execpb.NewExecutorServiceClient(executorcliConn)
	// api := newAPI(brokerapp.NewApp(cfg.Log, cfg.BrokerBus, execcli, authcli), cfg.Log)
	// api.execcli = execcli
	// api.authcli = authcli

	app.Use(mid.TraceIdMiddleware())

	// cfg.Log.Errorc("started serving ")
	app.Handle("GET", "/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})

}
