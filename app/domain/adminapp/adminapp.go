package adminapp

import (
	"context"
	"fmt"
	"time"

	"github.com/bentenison/microservice/api/domain/admin-api/grpc/proto/admin"
	"github.com/bentenison/microservice/business/domain/adminbus"
	"github.com/bentenison/microservice/foundation/async/kafka/kafkaconsumer"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
	"go.opentelemetry.io/otel/attribute"
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

func (a *App) CompleteChallenge(ctx context.Context, req *admin.CompleteChallengeRequest) (*admin.CompleteChallengeResponse, error) {

	return nil, nil
}
func (a *App) CompleteQuestion(ctx context.Context, req *admin.CompleteQuestionRequest) (*admin.CompleteQuestionResponse, error) {
	_, span := otel.AddSpan(ctx, "admin.CompleteQuestion", attribute.String("RPC", fmt.Sprintf("{started_at:%d}", time.Now().Unix())))
	defer span.End()
	var payload adminbus.UpdatePayload
	payload.ChallengeId = req.ChallengeId
	payload.QuestionId = req.QuestionId
	payload.IsCorrect = req.IsCorrect
	payload.CodeQuality = req.CodeQuality
	payload.IsChallenge = req.IsChallenge
	payload.Language = req.Language
	payload.UserId = req.UserId
	err := a.adminbus.MarkQuestionCompletion(ctx, payload)
	if err != nil {
		a.logger.Errorc(ctx, "error in MarkQuestionCompletion:", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	return &admin.CompleteQuestionResponse{Success: true}, nil
	// return nil, nil
}
func (a *App) UpdateUserMetrics(ctx context.Context, req *admin.UpdateUserMetricsRequest) (*admin.UpdateUserMetricsResponse, error) {
	var payload adminbus.UpdatePayload
	payload.ChallengeId = req.ChallengeId
	payload.QuestionId = req.QuestionId
	payload.IsCorrect = req.IsCorrect
	payload.CodeQuality = req.CodeQuality
	payload.IsChallenge = req.IsChallenge
	payload.Language = req.Language
	payload.UserId = req.UserId
	err := a.adminbus.UpdateUserMetrics(ctx, payload)
	if err != nil {
		a.logger.Errorc(ctx, "error in UpdateUserMetrics", map[string]interface{}{
			"error": err.Error(),
		})
		return &admin.UpdateUserMetricsResponse{Success: false}, err
	}
	return &admin.UpdateUserMetricsResponse{Success: true}, nil
}
func (a *App) UpdateUserPerformance(ctx context.Context, req *admin.UpdateUserPerformanceRequest) (*admin.UpdateUserPerformanceResponse, error) {
	var payload adminbus.UpdatePayload
	payload.ChallengeId = req.ChallengeId
	payload.QuestionId = req.QuestionId
	payload.IsCorrect = req.IsCorrect
	payload.CodeQuality = req.CodeQuality
	payload.IsChallenge = req.IsChallenge
	payload.Language = req.Language
	payload.UserId = req.UserId
	err := a.adminbus.UpdateUserStats(ctx, payload)
	if err != nil {
		a.logger.Errorc(ctx, "error in UpdateUserStats", map[string]interface{}{
			"error": err.Error(),
		})
		return &admin.UpdateUserPerformanceResponse{Success: false}, err
	}
	return &admin.UpdateUserPerformanceResponse{Success: true}, nil
}
func (a *App) AddSubmissionStats(ctx context.Context, req *admin.AddSubmissionStatsRequest) (*admin.AddSubmissionStatsResponse, error) {
	return nil, nil
}
