package adminapi

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bentenison/microservice/api/domain/admin-api/grpc/proto/admin"
	"github.com/bentenison/microservice/app/domain/adminapp"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	adminapp *adminapp.App
	logger   *logger.CustomLogger
	admin.UnimplementedAdminServiceServer
	// authcli   authpb.AuthServiceClient
	// execcli   execpb.ExecutorServiceClient
}

func newAPI(adminApp *adminapp.App, logger *logger.CustomLogger) *api {
	return &api{
		adminapp: adminApp,
		logger:   logger,
	}
}

func (api *api) createChallenge(c *gin.Context) {
	var user adminapp.User
	if err := c.Bind(&user); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	if user.UserID == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "userId is a required parameter",
		})
		c.Error(fmt.Errorf("userId is required")).SetMeta(http.StatusExpectationFailed)
		return
	}
	// submission := toAppSubmission(submissionPayload)
	challenge, err := api.adminapp.CreateChallenge(c.Request.Context(), adminapp.ToBusUser(user), user.SelectedLanguage)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, challenge)
}
func (api *api) prepareChallenge(c *gin.Context) {
	var user adminapp.User
	if err := c.Bind(&user); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	if user.UserID == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "userId is a required parameter",
		})
		c.Error(fmt.Errorf("userId is required")).SetMeta(http.StatusExpectationFailed)
		return
	}
	// submission := toAppSubmission(submissionPayload)
	err := api.adminapp.AddPreRequisites(c.Request.Context(), adminapp.ToBusUser(user), user.SelectedLanguage)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "OK")
}
func (api *api) fetchChallengeQuestions(c *gin.Context) {
	var ids []string
	if err := c.Bind(&ids); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	if len(ids) <= 0 {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "no questionIds present",
		})
		c.Error(fmt.Errorf("no questionIds present")).SetMeta(http.StatusExpectationFailed)
		return
	}
	// submission := toAppSubmission(submissionPayload)
	quests, err := api.adminapp.FetchAllQuestionsByIds(c.Request.Context(), ids)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, quests)
}
func changeQuestion(c *gin.Context) {
	// get the id of the question and challengeID in which the question to be chnaged
	// query the qc_question with the difficulty and questionid should not be in challenge questionids and select a random question
	// log that user has chnaged his question which will add a penalty of 5 points and a user can only change his question 3 times a day
	// update the questions array of the challenges collection and replace the previous id question with the newly created question
}
func saveAndShareSnippet(c *gin.Context) {
	// save snippet by userid to collection
	// return the url with id to the user
	// when user loads the url then get that code and load it to editor with language
}

func (api *api) CompleteChallenge(context.Context, *admin.CompleteChallengeRequest) (*admin.CompleteChallengeResponse, error) {
	api.logger.Errorc(context.TODO(), "completechallenge called", map[string]interface{}{})
	return nil, nil
}
func (api *api) CompleteQuestion(context.Context, *admin.CompleteQuestionRequest) (*admin.CompleteQuestionResponse, error) {
	log.Println("completeQuestion Called by RPC")
	return nil, nil
}
func (api *api) UpdateUserMetrics(context.Context, *admin.UpdateUserMetricsRequest) (*admin.UpdateUserMetricsResponse, error) {
	return nil, nil
}
func (api *api) UpdateUserPerformance(context.Context, *admin.UpdateUserPerformanceRequest) (*admin.UpdateUserPerformanceResponse, error) {
	return nil, nil
}
func (api *api) AddSubmissionStats(context.Context, *admin.AddSubmissionStatsRequest) (*admin.AddSubmissionStatsResponse, error) {
	return nil, nil
}
