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

func (b *Business) AuthenticateUser(username, password string) {
	// check if user exists in db
	// check if password hash matches
	// add session
	// return user token with role info

}
func (b *Business) AuthorizeUser(token string) {
	// get token
	// decrypt tokken
	// if token valid get userid
	// check session

}
