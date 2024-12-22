package adminapi

import (
	"context"
	"net/http"
	"time"

	"github.com/bentenison/microservice/api/domain/admin-api/grpc/proto/admin"
	"github.com/bentenison/microservice/api/sdk/grpc/rpcserver"
	"github.com/bentenison/microservice/app/domain/adminapp"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/adminbus"
	"github.com/bentenison/microservice/foundation/async/kafka/kafkaconsumer"
	"github.com/bentenison/microservice/foundation/async/rabbit/rabbitconsumer"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Log            *logger.CustomLogger
	AdminBus       *adminbus.Business
	Tracer         *trace.TracerProvider
	AppConfig      *conf.Config
	KafkaConsumer  *kafkaconsumer.Consumer
	RabbitConsumer *rabbitconsumer.Consumer
}

func Routes(app *web.App, cfg Config) {
	// authcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.AuthGRPCPort)
	// authcli := authpb.NewAuthServiceClient(authcliConn)
	// executorcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.GRPCPort)
	// execcli := execpb.NewExecutorServiceClient(executorcliConn)

	api := newAPI(adminapp.NewApp(cfg.Log, cfg.AdminBus, cfg.KafkaConsumer), cfg.Log)
	// api.execcli = execcli
	// api.authcli = authcli
	go RunGRPCServer(cfg.AppConfig.AdminGRPCPort, cfg.Log, api, cfg.Tracer)
	app.Use(mid.TraceIdMiddleware(), mid.Otel(cfg.Tracer.Tracer("ADMIN")), mid.ErrorMiddleware())
	app.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	// cfg.Log.Errorc("started serving ")
	app.Handle("GET", "/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})
	app.Handle("GET", "/leaderboard/:level", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})
	app.Handle("GET", "/predict-rank", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})
	app.Handle("GET", "/adaptive-challenge", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})
	app.Handle("GET", "/submit-score", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})
	app.Handle("GET", "/user-progress/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})
	app.Handle("GET", "/update-progress", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "running",
			"time":    time.Now(),
		})
	})
	app.Handle("POST", "/create-challenge", api.createChallenge)
	app.Handle("POST", "/prepare-challenge", api.prepareChallenge)
	app.Handle("POST", "/fetch-questions", api.fetchChallengeQuestions)
	app.Handle("POST", "/change-question", api.fetchChallengeQuestions)

}
func RunGRPCServer(GRPCPort string, log *logger.CustomLogger, api *api, tp *trace.TracerProvider) {
	grpcSrv, listner := rpcserver.CreateServer(GRPCPort, log, tp)
	defer listner.Close()
	// go func() {
	log.Infoc(context.TODO(), "startup grpc v1 server started", map[string]interface{}{
		"port": GRPCPort,
	})
	// executorServer := executorapi.NewExecutorServer(log)
	admin.RegisterAdminServiceServer(grpcSrv, api)
	reflection.Register(grpcSrv)
	if err := grpcSrv.Serve(listner); err != nil {
		log.Errorc(context.TODO(), "error occured while listning for grpc traffic", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// }()
}
