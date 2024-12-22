package creatorapp

import (
	"context"
	"time"

	"github.com/bentenison/microservice/business/domain/creatorbus"
	"github.com/bentenison/microservice/business/sdk/page"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	creatorbus *creatorbus.Business
	logger     *logger.CustomLogger
	tracer     trace.Tracer
	// authcli   authpb.AuthServiceClient
	// execcli   execpb.ExecutorServiceClient
}

func NewApp(logger *logger.CustomLogger, bus *creatorbus.Business) *App {
	return &App{logger: logger, creatorbus: bus}
}
func (a *App) AddNewQuestions(ctx context.Context, qts []Question) ([]interface{}, error) {
	now := time.Now()
	for i := 0; i < len(qts); i++ {
		genId := uuid.NewString()
		qts[i].QuestionId = genId
		qts[i].Answer.ID = genId
		qts[i].Answer.CreatedAt = &now
		qts[i].Answer.UpdatedAt = &now
		qts[i].IsQC = false
	}
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
func (a *App) Query(ctx context.Context, page page.Page, qp QueryParams) (creatorbus.QueryResult, error) {
	filter, err := parseFilter(qp)
	if err != nil {
		a.logger.Errorc(ctx, "error while parsing filter", map[string]interface{}{
			"error": err.Error(),
		})
		return creatorbus.QueryResult{}, err
	}
	return a.creatorbus.Query(ctx, filter, page)
}
func (a *App) GetAllProgrammingConcepts(ctx context.Context) ([]creatorbus.LanguageConcept, error) {
	res, err := a.creatorbus.GetAllLanguageConcepts(ctx)
	if err != nil {
		a.logger.Errorc(ctx, "error in GetAllProgrammingConcepts", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	return res, nil
}
