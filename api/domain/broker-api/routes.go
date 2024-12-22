package brokerapi

import (
	"context"
	"log"

	"github.com/bentenison/microservice/api/domain/broker-api/grpc/adminclient/proto/admClient"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto/authCli"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto/execClient"
	brokerapp "github.com/bentenison/microservice/app/domain/broker-app"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Log            *logger.CustomLogger
	BrokerBus      *brokerbus.Business
	Tracer         *trace.TracerProvider
	AppConfig      *conf.Config
	AuthClient     authCli.AuthServiceClient
	ExecutorClient execClient.ExecutorServiceClient
	AdminClient    admClient.AdminServiceClient
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(brokerapp.NewApp(cfg.Log, cfg.BrokerBus, cfg.Tracer, cfg.AdminClient, cfg.ExecutorClient, cfg.AuthClient), cfg.Log, cfg.AuthClient, cfg.ExecutorClient, cfg.AdminClient)
	// api.execcli = execcli
	// api.authcli = authcli
	// api.admincli = admincli
	app.Use(mid.TraceIdMiddleware(), mid.Otel(cfg.Tracer.Tracer("")), mid.ErrorMiddleware())

	// cfg.Log.Errorc("started serving ")
	app.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	app.Handle("POST", "/broker/submission", api.newSubmissionHandler)
	app.Handle("POST", "/broker/run", api.codeRunHandler)
	app.Handle("POST", "/broker/qcquestion", api.qcQuestion)
	app.Handle("POST", "/broker/batchProcess", api.newSubmissionHandler)
	app.Handle("POST", "/broker/authenticate", api.authenticateHandler)
	app.Handle("POST", "/broker/authorize", api.authorizeHandler)
	app.Handle("POST", "/broker/create", api.createUserHandler)
	app.Handle("POST", "/broker/getllquestions", api.getAllQuestionsHandler)
	app.Handle("POST", "/broker/getquestion/:id", api.getQuestionHandler)
	app.Handle("POST", "/broker/getanswer/:id", api.getAnswerHandler)
	app.Handle("POST", "/broker/getallanswer", api.getAllAnswersHandler)
	app.Handle("GET", "/broker/getlanguages", api.getallLanguages)
	app.Handle("GET", "/broker/getallowedlanguages", api.newSubmissionHandler)
	app.Handle("GET", "/broker/conf", api.newSubmissionHandler)
	app.Handle("GET", "/broker/updateconf", api.newSubmissionHandler)
	app.Handle("GET", "/broker/gettemplates", api.getAllQuestionTemplates)
	app.Handle("GET", "/broker/addsnippet", api.createSnippet)
	app.Handle("GET", "/broker/getsnippet/:id", api.getSnippetById)
	app.Handle("GET", "/broker/getAllsnippets", api.getAllSnippetsByUser)
	app.Handle("POST", "/broker/formatCode", api.formatCode)
	app.Handle("GET", "/broker/loaddbQuest/:id", api.loadDbQuestion)
	app.Handle("GET", "/text-rpc", func(ctx *gin.Context) {
		res, err := api.admincli.CompleteQuestion(context.TODO(), &admClient.CompleteQuestionRequest{})
		if err != nil {
			log.Println(err)
		}
		_ = res
	})
}
