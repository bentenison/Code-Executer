package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bentenison/microservice/api/cmd/service/exam-service/build/all"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/sdk/mongodb"
	"github.com/bentenison/microservice/business/sdk/sqldb"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

var apiType = "all"

func main() {
	//initialize the application logger
	log := logger.NewCustomLogger(map[string]interface{}{
		"service": "exam-service",
		"env":     "production",
		"build":   "1.0.0",
	})
	// load congigurations
	config, err := conf.LoadConfig()
	if err != nil {
		log.Errorc(context.TODO(), "error while loading conf", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// // -------------------------------------------------------------------------
	// // INITIALIZE TRACER OTEL
	cfg := otel.Config{
		Host:        config.TracerHost,
		Probability: config.TracerProb,
		ServiceName: "BROKER",
	}
	trace, err := otel.NewTracer(cfg)
	if err != nil {
		log.Errorc(context.TODO(), "error while initializing tracer", map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer func() {
		otel.ShutDownTracer(trace)
	}()
	log.Infoc(context.TODO(), "config", map[string]interface{}{"config": config})
	if err := run(log, trace, config); err != nil {
		log.Errorc(context.TODO(), "error while running server", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// log.Errorc("error while loading conf", map[string]interface{}{
	// 	"error": "error",
	// })
}
func run(log *logger.CustomLogger, tracer *trace.TracerProvider, cfg *conf.Config) error {
	//starting database connection
	db, err := sqldb.Open(sqldb.Config{
		User:         cfg.User,
		Password:     cfg.Password,
		Host:         cfg.Host,
		Name:         cfg.DBName,
		MaxIdleConns: cfg.MaxIdleConns,
		MaxOpenConns: cfg.MaxOpenConns,
	})
	if err != nil {
		// log.Errorc("error connecting to DB")
		return err
	}

	defer db.Close()

	mongo, err := mongodb.InitializeMongo(mongodb.Config{
		Host:     cfg.MongoHost,
		Port:     cfg.MongoPort,
		Username: cfg.MongoUser,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	ds := mux.DataSource{
		MGO: mongo,
		SQL: db,
	}
	// go func() {
	// 	log.Infoc(context.TODO(), "startup debug v1 server started", map[string]interface{}{
	// 		"port": cfg.DebugPort,
	// 	})

	// 	if err := http.ListenAndServe(cfg.DebugPort, debug.Mux()); err != nil {
	// 		log.Errorc(context.TODO(), "error occured while listning for traffic", map[string]interface{}{
	// 			"error": err.Error(),
	// 		})
	// 	}
	// }()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := mux.Config{
		Build: "develop",
		Log:   log,
		DB:    ds,
		// Tracer: tracer,
	}
	app := mux.WebAPI(cfgMux, buildRoutes())
	api := http.Server{
		Addr:    cfg.ExamAPIPort,
		Handler: app,
		// ReadTimeout:  cfg.Web.ReadTimeout,
		// WriteTimeout: cfg.,
		// IdleTimeout:  cfg.Web.IdleTimeout,
		// ErrorLog: lo,
	}

	serverErrors := make(chan error, 1)
	ctx := context.Background()
	go func() {
		log.Infoc(context.TODO(), "api router started", map[string]interface{}{
			"port": cfg.ExamAPIPort,
		})
		serverErrors <- api.ListenAndServe()
	}()

	// // -------------------------------------------------------------------------
	// // Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infoc(context.TODO(), "shutdown started", map[string]interface{}{
			"signal": sig,
		})
		defer log.Infoc(context.TODO(), "shutdown completed", map[string]interface{}{})

		ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}

func buildRoutes() mux.RouteAdder {
	switch apiType {
	case "all":
		return all.Routes()

	}
	return all.Routes()
}
