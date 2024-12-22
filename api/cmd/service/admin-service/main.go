package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bentenison/microservice/api/cmd/service/admin-service/build/all"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	essearch "github.com/bentenison/microservice/business/sdk/elastricsearch"
	"github.com/bentenison/microservice/business/sdk/mongodb"
	"github.com/bentenison/microservice/business/sdk/redisdb"
	"github.com/bentenison/microservice/business/sdk/sqldb"
	"github.com/bentenison/microservice/foundation/async/rabbit/rabbitconsumer"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/monitoring"
	"github.com/bentenison/microservice/foundation/otel"
	"github.com/elastic/go-elasticsearch/v8"
	"go.opentelemetry.io/otel/sdk/trace"
)

const apiType = "all"
const userPerformanceMapping = `{
	"mappings": {
		"properties": {
			"user_id": { "type": "keyword" },
			"accuracy": { "type": "float" },
			"speed_avg": { "type": "float" },
			"penalty_points": { "type": "integer" },
			"question_id": { "type": "keyword" },
			"language": { "type": "keyword" },
			"rank": { "type": "integer" },
			"created_at": { "type": "date" }
		}
	}
}`

const challengeDataMapping = `{
	"mappings": {
		"properties": {
			"challenge_id": { "type": "keyword" },
			"difficulty": { "type": "integer" },
			"score": { "type": "integer" },
			"accuracy": { "type": "float" },
			"speed_avg": { "type": "float" },
			"language": { "type": "keyword" },
			"penalty_points": { "type": "integer" },
			"execution_time": { "type": "integer" }
		}
	}
}`

const codeExecutionStatsMapping = `{
	"mappings": {
		"properties": {
			"execution_time": { "type": "float" },
			"memory_usage": { "type": "long" },
			"total_memory": { "type": "long" },
			"cpu_usage": { "type": "long" },
			"memory_percentage": { "type": "float" },
			"created_at": { "type": "date" },
			"updated_at": { "type": "date" },
			"id": { "type": "keyword" },
			"user_id": { "type": "keyword" },
			"language_id": { "type": "keyword" },
			"status": { "type": "keyword" },
			"error_message": { "type": "text" },
			"code_snippet": { "type": "text" },
			"container_id": { "type": "keyword" }
		}
	}
}`

func main() {
	log := logger.NewCustomLogger(map[string]interface{}{
		"service": "admin-service",
		"env":     "production",
		"build":   "1.0.0",
	})
	// load configurations
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
		ServiceName: "ADMIN",
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
	if err := run(log, config, trace); err != nil {
		log.Errorc(context.TODO(), "error while running server", map[string]interface{}{
			"error": err.Error(),
		})
	}
}
func run(log *logger.CustomLogger, cfg *conf.Config, tracer *trace.TracerProvider) error {
	//starting sql database connection
	db, err := sqldb.Open(sqldb.Config{
		User:         cfg.User,
		Password:     cfg.Password,
		Host:         cfg.Host,
		Name:         cfg.DBName,
		MaxIdleConns: cfg.MaxIdleConns,
		MaxOpenConns: cfg.MaxOpenConns,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}

	defer db.Close()
	// starting mongio db connection

	mongo, err := mongodb.InitializeMongo(mongodb.Config{
		Username:    cfg.MongoUser,
		Password:    cfg.MongoPassword,
		AuthDB:      cfg.MongoAuth,
		Host:        cfg.MongoHost,
		Port:        cfg.MongoPort,
		DBName:      cfg.MongoDbName,
		AllowDirect: cfg.AllowDirect,
	})
	if err != nil {
		return err
	}
	rdb, err := redisdb.OpenRDB(redisdb.Config{})
	if err != nil {
		return err
	}
	es, err := essearch.InitElasticsearch()
	if err != nil {
		return err
	}
	mappings := map[string]string{
		"user_performance":     userPerformanceMapping,
		"challenge_data":       challengeDataMapping,
		"code_execution_stats": codeExecutionStatsMapping,
	}
	for k, v := range mappings {
		err := createIndexIfNotExists(es, log, k, v)
		if err != nil {
			log.Errorc(context.TODO(), "error creating index", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
	}
	log.Infoc(context.TODO(), "connected to ES and indexes created", map[string]interface{}{})
	ds := mux.DataSource{
		MGO: mongo,
		SQL: db,
		RDB: rdb,
		ES:  es,
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
	//starts kafka consumer
	// broker := []string{"localhost:9092"}
	// groupId := []string{"es_analytics", "performnace"}
	// topics := []string{"user_performance", "challenge_data", "code_execution_stats"}
	// consumer, err := kafkaconsumer.NewConsumer(broker, groupId, topics)
	// if err != nil {
	// 	log.Errorc(context.TODO(), "error occured while listning kafka topics", map[string]interface{}{
	// 		"error": err.Error(),
	// 	})
	// }
	// _ = consumer
	queues := strings.Split(cfg.RabbitQueues, ",")
	consumer, err := rabbitconsumer.NewConsumer(cfg.RabbitURL, queues, log, es)
	if err != nil {
		log.Errorc(context.TODO(), "error occured while starting rabbit consumer", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	monitoring.StartMetricsForService("Admin")
	monitoring.CollectSystemMetrics()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := mux.Config{
		Build:          "develop",
		Log:            log,
		DB:             ds,
		AppConfig:      cfg,
		Tracer:         tracer,
		RabbitConsumer: consumer,
		// KafkaConsumer: consumer,
	}
	app := mux.WebAPI(cfgMux, buildRoutes())
	api := http.Server{
		Addr:    cfg.AdminAPIPort,
		Handler: app,
		// ReadTimeout:  cfg.Web.ReadTimeout,
		// WriteTimeout: cfg.,
		// IdleTimeout:  cfg.Web.IdleTimeout,
		// ErrorLog: lo,
	}

	serverErrors := make(chan error, 1)
	ctx := context.Background()
	go func() {
		log.Infoc(context.TODO(), "admin-api router started", map[string]interface{}{
			"port": cfg.AdminAPIPort,
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

func createIndexIfNotExists(es *elasticsearch.Client, log *logger.CustomLogger, index string, mapping string) error {
	// Check if index exists
	res, err := es.Indices.Exists([]string{index})
	if err != nil {
		return fmt.Errorf("failed to check if index exists: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		// Create the index if it doesn't exist
		createIndexRes, err := es.Indices.Create(index, es.Indices.Create.WithBody(strings.NewReader(mapping)))
		if err != nil {
			return fmt.Errorf("failed to create index: %w", err)
		}
		defer createIndexRes.Body.Close()

		if createIndexRes.IsError() {
			return fmt.Errorf("error creating index: %s", createIndexRes.String())
		}
		log.Errorc(context.TODO(), "index created successfully", map[string]interface{}{
			"index": fmt.Sprintf("Index %s created.", index),
		})
	} else {
		log.Errorc(context.TODO(), "index already exists.", map[string]interface{}{
			"index": fmt.Sprintf("Index %s already exists.", index),
		})
	}

	return nil
}
