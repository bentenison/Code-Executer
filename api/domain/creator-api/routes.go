package creatorapi

import (
	creatorapp "github.com/bentenison/microservice/app/domain/creator-app"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/creatorbus"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	Log        *logger.CustomLogger
	CreatorBus *creatorbus.Business
	Tracer     *trace.TracerProvider
	AppConfig  *conf.Config
}

func Routes(app *web.App, cfg Config) {
	// authcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.AuthGRPCPort)
	// authcli := authpb.NewAuthServiceClient(authcliConn)
	// executorcliConn := rpcserver.CreateClient(cfg.Log, cfg.AppConfig.GRPCPort)
	// execcli := execpb.NewExecutorServiceClient(executorcliConn)
	api := newAPI(cfg.Log, creatorapp.NewApp(cfg.Log, cfg.CreatorBus))
	// api.execcli = execcli
	// api.authcli = authcli

	app.Use(mid.TraceIdMiddleware())

	// cfg.Log.Errorc("started serving ")
	app.Handle("POST", "/creator/addquestions", api.createNewQuestions)
	app.Handle("POST", "/creator/qcquestion", api.qcQuestion)
	app.Handle("POST", "/creator/getallquestions", api.getAllQuestions)
	app.Handle("GET", "/creator/getquestiontag/:tag", api.getQuestionByTag)
	app.Handle("GET", "/creator/getsingleQuestion/:id", api.getSingleQuestion)
	app.Handle("GET", "/creator/getquestion/:lang", api.getQuestionsBylang)
	app.Handle("POST", "/creator/deleteQuestions", api.deleteSelectedQuestion)
	app.Handle("GET", "/creator/query", api.query)
	app.Handle("GET", "/creator/languageConcepts", api.getAllLanguageConcepts)
	// app.Handle("POST", "/creator/deleteselected", api.createNewQuestions)
}
