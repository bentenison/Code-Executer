package brokerbus

import (
	"context"
	"os"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
)

type Storer interface {
	GetQuestionTemplate(ctx context.Context) (Template, error)
}

type Business struct {
	log      *logger.CustomLogger
	db       mux.DataSource
	delegate *delegate.Delegate
	storer   Storer
}

func NewBusiness(logger *logger.CustomLogger, ds mux.DataSource, delegate *delegate.Delegate, storer Storer) *Business {
	return &Business{
		log:      logger,
		db:       ds,
		delegate: delegate,
		storer:   storer,
	}
}

func (b *Business) GetQuestionTemplate(ctx context.Context) (Template, error) {
	template, err := b.storer.GetQuestionTemplate(ctx)
	if err != nil {
		b.log.Error("error while getting template", map[string]interface{}{
			"error": err,
			"trace": mid.GetTraceId(ctx),
		})
		return Template{}, err
	}
	f, err := os.OpenFile("code.py", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		b.log.Error("error while creating file", map[string]interface{}{
			"error": err,
			"trace": mid.GetTraceId(ctx),
		})
		return template, err
	}
	defer f.Close()
	f.WriteString(template.Template)
	return template, err
}
