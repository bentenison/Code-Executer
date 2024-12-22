package brokerapp

import (
	"context"

	"github.com/bentenison/microservice/api/domain/broker-api/grpc/adminclient/proto/admClient"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto/authCli"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto/execClient"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	"go.mongodb.org/mongo-driver/mongo"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	brokerbus *brokerbus.Business
	logger    *logger.CustomLogger
	tracer    trace.Tracer
	authcli   authCli.AuthServiceClient
	execcli   execClient.ExecutorServiceClient
	admincli  admClient.AdminServiceClient
}

func NewApp(logger *logger.CustomLogger, bus *brokerbus.Business, tp *sdktrace.TracerProvider, admincli admClient.AdminServiceClient, execcli execClient.ExecutorServiceClient, authcli authCli.AuthServiceClient) *App {
	return &App{logger: logger, brokerbus: bus, tracer: tp.Tracer(""), authcli: authcli, execcli: execcli, admincli: admincli}
}

func (a *App) HandleSubmisson(ctx context.Context, submission Submission) (*execClient.ExecutionResponse, error) {
	return a.brokerbus.HandleSubmissonService(ctx, brokerbus.Submission(submission), a.authcli, a.execcli)
}
func (a *App) HandleCodeRun(ctx context.Context, submission Submission) (*execClient.ExecutionResponse, error) {
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
func (a *App) GetAllLanguages(ctx context.Context) ([]*brokerbus.Language, error) {
	return a.brokerbus.GetAllAllowedLanguages(ctx)
}
func (a *App) HandleQCQuestion(ctx context.Context, q Question) (*execClient.ExecutionResponse, error) {
	return a.brokerbus.HandleQcService(ctx, toBusQuestion(q), a.authcli, a.execcli)
}
func (a *App) HandleGetAllTemplates(ctx context.Context) ([]brokerbus.Question, error) {
	return a.brokerbus.GetAllQuestTemplates(ctx)
}
func (a *App) HandleCreateSnippet(ctx context.Context, snippet CodeSnippet) (*mongo.InsertOneResult, error) {
	brokerBusCodeSnippet := brokerbus.CodeSnippet(snippet)
	return a.brokerbus.CreateCodeSnippet(ctx, &brokerBusCodeSnippet)
}
func (a *App) HandleGetAllSnippets(ctx context.Context, userId string) ([]brokerbus.CodeSnippet, error) {
	return a.brokerbus.GetAllSnippetsByUser(ctx, userId)
}
func (a *App) HandleGetSnippetById(ctx context.Context, id string) (*brokerbus.CodeSnippet, error) {
	return a.brokerbus.GetSnippetById(ctx, id)
}
func (a *App) FormatCode(ctx context.Context, payload FormatterRequest) (*brokerbus.FormatterResponse, error) {
	return a.brokerbus.FormatCode(ctx, brokerbus.FormatterRequest(payload))
}
func (a *App) LoadDBQuest(ctx context.Context, questId string) (*brokerbus.SQLQuestion, error) {
	return a.brokerbus.GetDBQuestById(ctx, questId)
}
