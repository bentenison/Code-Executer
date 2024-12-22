package creatorbus

import (
	"context"

	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/business/sdk/page"
	"github.com/bentenison/microservice/foundation/logger"
)

type Storer interface {
	DeleteAnswer(ctx context.Context, id string) (int64, error)
	DeleteAnswers(ctx context.Context, ids []string) (int64, error)
	DeleteQuestion(ctx context.Context, id string) (int64, error)
	DeleteQuestions(ctx context.Context, ids []string) (int64, error)
	GetAllAnswersDAO(ctx context.Context) ([]Answer, error)
	GetAllQuestionsDAO(ctx context.Context) ([]Question, error)
	GetAnswerById(ctx context.Context, id string) (Answer, error)
	GetSingleQuestion(ctx context.Context, id string) (Question, error)
	AddQuestions(ctx context.Context, qts []Question) ([]interface{}, error)
	AddQCQuestions(ctx context.Context, qts []Question) ([]interface{}, error)
	AddAnswer(ctx context.Context, ans []Answer) ([]interface{}, error)
	SearchQuestionByTag(ctx context.Context, tag string) ([]Question, error)
	SearchQuestionByLang(ctx context.Context, lang string) ([]Question, error)
	Query(ctx context.Context, filter QueryFilter, page page.Page) (QueryResult, error)
	GetAllLanguageConcepts(ctx context.Context) ([]LanguageConcept, error)
}

type Business struct {
	log      *logger.CustomLogger
	delegate *delegate.Delegate
	storer   Storer
}

func NewBusiness(logger *logger.CustomLogger, delegate *delegate.Delegate, storer Storer) *Business {
	return &Business{
		log:      logger,
		delegate: delegate,
		storer:   storer,
	}
}

func (b *Business) AddNewQuestions(ctx context.Context, qts []Question) ([]interface{}, error) {
	return b.storer.AddQCQuestions(ctx, qts)
}

func (b *Business) GetSingleQuestion(ctx context.Context, id string) (Question, error) {
	return b.storer.GetSingleQuestion(ctx, id)
}

func (b *Business) DeleteSelectedQuestions(ctx context.Context, ids []string) (int64, error) {
	return b.storer.DeleteQuestions(ctx, ids)
}

// func (b *Business) QCQuestion() {

// }
func (b *Business) GetAllQuestionsDAO(ctx context.Context) ([]Question, error) {
	return b.storer.GetAllQuestionsDAO(ctx)
}
func (b *Business) GetQuestionsByTag(ctx context.Context, tag string) ([]Question, error) {
	return b.storer.SearchQuestionByTag(ctx, tag)
}
func (b *Business) GetQuestionsByLang(ctx context.Context, lang string) ([]Question, error) {
	return b.storer.SearchQuestionByLang(ctx, lang)
}
func (b *Business) Query(ctx context.Context, filter QueryFilter, page page.Page) (QueryResult, error) {
	return b.storer.Query(ctx, filter, page)
}
func (b *Business) GetAllLanguageConcepts(ctx context.Context) ([]LanguageConcept, error) {
	return b.storer.GetAllLanguageConcepts(ctx)
}
