package adminapp

import (
	"context"

	"github.com/bentenison/microservice/business/domain/adminbus"
	"github.com/bentenison/microservice/foundation/async/kafka/kafkaconsumer"
	"github.com/bentenison/microservice/foundation/logger"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	adminbus      *adminbus.Business
	logger        *logger.CustomLogger
	tracer        trace.Tracer
	kafkaconsumer *kafkaconsumer.Consumer
}

func NewApp(logger *logger.CustomLogger, bus *adminbus.Business, consumer *kafkaconsumer.Consumer) *App {
	return &App{logger: logger, adminbus: bus, kafkaconsumer: consumer}
}
func (a *App) AddPreRequisites(ctx context.Context, user adminbus.User, language string) error {
	err := a.adminbus.AddPreRequisites(ctx, user, language)
	if err != nil {
		return err
	}
	return nil
}
func (a *App) CreateChallenge(ctx context.Context, user adminbus.User, language string) (adminbus.Challenge, error) {
	res, err := a.adminbus.CreateChallengeService(ctx, user, language)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (a *App) FetchAllQuestionsByIds(ctx context.Context, ids []string) ([]adminbus.CodingQuestion, error) {
	res, err := a.adminbus.FetchQuestionsByIds(ctx, ids)
	if err != nil {
		return res, err
	}
	return res, nil
}
