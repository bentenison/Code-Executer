package brokerdb

import (
	"context"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type Store struct {
	ds     mux.DataSource
	logger *logger.CustomLogger
}

func NewStore(ds mux.DataSource, logger *logger.CustomLogger) *Store {
	return &Store{
		ds:     ds,
		logger: logger,
	}
}

func (s *Store) GetQuestionTemplate(ctx context.Context) (brokerbus.Template, error) {
	template := brokerbus.Template{}
	res := s.ds.MGO.Collection("question_template").FindOne(ctx, bson.M{})
	if res.Err() != nil {
		return template, res.Err()
	}
	err := res.Decode(&template)
	if err != nil {
		return template, err
	}
	return template, nil
}
