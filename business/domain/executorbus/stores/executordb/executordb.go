package executordb

import (
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/foundation/logger"
)

type Store struct {
	log *logger.CustomLogger
	db  mux.DataSource
}

func NewStore(log *logger.CustomLogger, ds mux.DataSource) *Store {
	return &Store{
		log: log,
		db:  ds,
	}
}
