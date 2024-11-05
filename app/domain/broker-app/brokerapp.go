package brokerapp

import (
	"context"

	authpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto"
	execpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	brokerbus *brokerbus.Business
	logger    *logger.CustomLogger
	tracer    trace.Tracer
	authcli   authpb.AuthServiceClient
	execcli   execpb.ExecutorServiceClient
}

func NewApp(logger *logger.CustomLogger, bus *brokerbus.Business, tp *tp.TracerProvider, execcli execpb.ExecutorServiceClient, authcli authpb.AuthServiceClient) *App {
	return &App{logger: logger, brokerbus: bus, tracer: tp.Tracer("BROKER"), authcli: authcli, execcli: execcli}
}

func (a *App) HandleSubmisson(ctx context.Context, submission Submission) (*execpb.ExecutionResponse, error) {
	return a.brokerbus.HandleSubmissonService(ctx, brokerbus.Submission(submission), a.authcli, a.execcli)
}
func (a *App) HandleCodeRun(ctx context.Context, submission Submission) (*execpb.ExecutionResponse, error) {
	return a.brokerbus.HandleCodeRun(ctx, brokerbus.Submission(submission), a.authcli, a.execcli)
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
func (a *App) GetAllQuestions(ctx context.Context) ([]brokerbus.Question, error) {
	return a.brokerbus.GetAllQuestions(ctx)
}
func (a *App) GetQuestionById(ctx context.Context, id string) (brokerbus.Question, error) {
	return a.brokerbus.GetSingleQuestion(ctx, id)
}
func (a *App) GetAllAnswers(ctx context.Context) ([]brokerbus.Answer, error) {
	return a.brokerbus.GetAllAnswers(ctx)
}
func (a *App) GetAnswerByQuestionId(ctx context.Context, id string) (brokerbus.Answer, error) {
	return a.brokerbus.GetAnswerByQuestion(ctx, id)
}
