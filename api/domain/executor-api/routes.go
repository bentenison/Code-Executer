package executorapi

import (
	"context"
	"net/http"

	pb "github.com/bentenison/microservice/api/domain/executor-api/grpc/proto"
	"github.com/bentenison/microservice/api/sdk/grpc/rpcserver"
	"github.com/bentenison/microservice/app/domain/executorapp"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/executorbus"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Log         *logger.CustomLogger
	Tracer      *trace.TracerProvider
	ExecutorBus *executorbus.Business
	AppConfig   *conf.Config
	// DockerClient *client.Client
}

func Routes(app *web.App, conf Config) {
	api := newAPI(executorapp.NewApp(conf.ExecutorBus, conf.Log, conf.Tracer), conf.Log)
	// ENABLE GRPC SERVER
	app.Use(mid.TraceIdMiddleware(), mid.Otel(conf.Tracer.Tracer("EXECUTOR")))
	go RunGRPCServer(conf.AppConfig.GRPCPort, conf.Log, api, conf.Tracer)
	app.Handle("POST", "executor/handlesubmission", api.handleSubmission)
	app.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	app.Handle("GET", "/check", func(c *gin.Context) {
		// ctx := c.Request.Context()
		// log.Println("tracer id", ctx.Value(otel.TraceIDKey))
		// log.Println("tracer id", ctx.Value("tracectx"))
		_, span := otel.AddSpan(c.Request.Context(), "HealthCheck")
		defer span.End()
		c.JSON(http.StatusOK, gin.H{"status": "running"})
	})
}

func RunGRPCServer(GRPCPort string, log *logger.CustomLogger, api *api, tp *trace.TracerProvider) {
	grpcSrv, listner := rpcserver.CreateServer(GRPCPort, log, tp)
	defer listner.Close()
	// go func() {
	log.Infoc(context.TODO(), "startup grpc v1 server started", map[string]interface{}{
		"port": GRPCPort,
	})
	// executorServer := executorapi.NewExecutorServer(log)
	pb.RegisterExecutorServiceServer(grpcSrv, api)
	reflection.Register(grpcSrv)
	if err := grpcSrv.Serve(listner); err != nil {
		log.Errorc(context.TODO(), "error occured while listning for grpc traffic", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// }()
}
