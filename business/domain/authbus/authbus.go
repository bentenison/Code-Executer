package authbus

import (
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
)

type Storer interface{}
type Business struct {
	log      *logger.CustomLogger
	delegate *delegate.Delegate
	ds       mux.DataSource
	storer   Storer
}

func NewBusiness(log *logger.CustomLogger, delegate *delegate.Delegate, ds mux.DataSource, storer Storer) *Business {
	return &Business{
		log:      log,
		delegate: delegate,
		ds:       ds,
		storer:   storer,
	}
}
