package exambus

import (
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
)

type Storer interface {
	// GetQuestionTemplate(ctx context.Context, id string) (Question, error)
	// AddSubmission(ctx context.Context, submission *Submission) (string, error)
	// AddExecutionStats(ctx context.Context, newStat *CodeExecutionStats) (string, error)
	// GetLanguages(ctx context.Context) ([]*Language, error)
	// GetAllQuestionsDAO(ctx context.Context) ([]Question, error)
	// GetAllAnswersDAO(ctx context.Context) ([]Answer, error)
	// GetAnswerById(ctx context.Context, id string) (Answer, error)
	// UpdateQCQuestion(ctx context.Context, id string) (*mongo.UpdateResult, error)
	// GetQuestionTemplates(ctx context.Context) ([]Question, error)
	// Get(ctx context.Context, key string, res any) error
	// Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error)
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
