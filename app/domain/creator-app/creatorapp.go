package creatorapp

import (
	"context"

	"github.com/bentenison/microservice/business/domain/creatorbus"
	"github.com/bentenison/microservice/foundation/logger"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	creatorbus *creatorbus.Business
	logger     *logger.CustomLogger
	tracer     trace.Tracer
	// authcli   authpb.AuthServiceClient
	// execcli   execpb.ExecutorServiceClient
}

func NewApp(logger *logger.CustomLogger, bus *creatorbus.Business, tp *tp.TracerProvider) *App {
	return &App{logger: logger, creatorbus: bus, tracer: tp.Tracer("Q-CREATOR")}
}
func (a *App) AddNewQuestions(ctx context.Context, qts []Question) ([]interface{}, error) {
	return a.creatorbus.AddNewQuestions(ctx, toBusQuestions(qts))
}

func (a *App) GetSingleQuestion(ctx context.Context, id string) (creatorbus.Question, error) {
	return a.creatorbus.GetSingleQuestion(ctx, id)
}

func (a *App) DeleteSelectedQuestions(ctx context.Context, ids []string) (int64, error) {
	return a.creatorbus.DeleteSelectedQuestions(ctx, ids)
}

func (a *App) QCQuestion() {

}
func (a *App) GetAllQuestionsDAO(ctx context.Context) ([]creatorbus.Question, error) {
	return a.creatorbus.GetAllQuestionsDAO(ctx)
}
func (a *App) GetQuestionsByTag(ctx context.Context, tag string) ([]creatorbus.Question, error) {
	return a.creatorbus.GetQuestionsByTag(ctx, tag)
}
func (a *App) GetQuestionsByLang(ctx context.Context, lang string) ([]creatorbus.Question, error) {
	return a.creatorbus.GetQuestionsByLang(ctx, lang)
}
