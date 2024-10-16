package brokerapp

import (
	"context"

	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	brokerbus *brokerbus.Business
	logger    *logger.CustomLogger
	tracer    trace.Tracer
}

func NewApp(logger *logger.CustomLogger, bus *brokerbus.Business, tp *tp.TracerProvider) *App {
	return &App{logger: logger, brokerbus: bus, tracer: tp.Tracer("BROKER")}
}

func (a *App) HandleSubmisson(ctx context.Context, submission Submission) (brokerbus.Question, error) {
	return a.brokerbus.HandleSubmissonService(ctx, brokerbus.Submission(submission))
}
func (a *App) Authenticate(ctx context.Context, up Credentials) (string, error) {
	return a.brokerbus.HandleAuthentication(ctx, up.Username, up.Password)
}

func (a *App) Authorize(ctx context.Context, token string) (brokerbus.Claims, error) {
	return a.brokerbus.HandleAuthorization(ctx, token)
}
func (a *App) CreateUser(ctx context.Context, up UserPayload) (string, error) {
	return a.brokerbus.HandleCreation(ctx, brokerbus.UserPayload(up))
}
