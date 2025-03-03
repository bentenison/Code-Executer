package brokerapi

import (
	"fmt"
	"net/http"

	"github.com/bentenison/microservice/api/domain/broker-api/grpc/adminclient/proto/admClient"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto/authCli"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto/execClient"
	brokerapp "github.com/bentenison/microservice/app/domain/broker-app"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	brokerapp *brokerapp.App
	logger    *logger.CustomLogger
	authcli   authCli.AuthServiceClient
	execcli   execClient.ExecutorServiceClient
	admincli  admClient.AdminServiceClient
}

func newAPI(brokerApp *brokerapp.App, logger *logger.CustomLogger, authcli authCli.AuthServiceClient, execcli execClient.ExecutorServiceClient, admincli admClient.AdminServiceClient) *api {
	return &api{
		brokerapp: brokerApp,
		logger:    logger,
		authcli:   authcli,
		execcli:   execcli,
		admincli:  admincli,
	}
}

func (api *api) newSubmissionHandler(c *gin.Context) {
	var submissionPayload SubmissionPayload
	if err := c.Bind(&submissionPayload); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	submission := toAppSubmission(submissionPayload)
	template, err := api.brokerapp.HandleSubmisson(c.Request.Context(), submission)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, template)
}
func (api *api) codeRunHandler(c *gin.Context) {
	var submissionPayload SubmissionPayload
	if err := c.Bind(&submissionPayload); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	submission := toAppSubmission(submissionPayload)
	template, err := api.brokerapp.HandleCodeRun(c.Request.Context(), submission)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, template)
}
func (api *api) authenticateHandler(c *gin.Context) {
	var cred brokerapp.Credentials
	if err := c.Bind(&cred); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	// submission := toAppSubmission(submissionPayload)
	token, err := api.brokerapp.Authenticate(c.Request.Context(), cred)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting token:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, token)
}
func (api *api) authorizeHandler(c *gin.Context) {
	var tkn token
	if err := c.Bind(&tkn); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}

	res, err := api.brokerapp.Authorize(c.Request.Context(), tkn.Token)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while authorizing user:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) createUserHandler(c *gin.Context) {
	var up brokerapp.UserPayload
	if err := c.Bind(&up); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	// submission := toAppSubmission(submissionPayload)
	template, err := api.brokerapp.CreateUser(c.Request.Context(), up)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, template)
}

func (api *api) authenticate(c *gin.Context) {
	var cred Credentials
	if err := c.Bind(&cred); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.authcli.Authenticate(c.Request.Context(), &authCli.LoginRequest{Username: cred.Username, Password: cred.Password})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in Authenticate GRPC API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) authorize(c *gin.Context) {
	var tkn token
	if err := c.Bind(&tkn); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.authcli.Authorize(c.Request.Context(), &authCli.AuthorizeRequest{Token: tkn.Token})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in Authorize GRPC API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) createUser(c *gin.Context) {
	var user UserPayload
	if err := c.Bind(&user); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.authcli.CreateAccount(c.Request.Context(), &authCli.CreateAccountRequest{Username: user.Username, Email: user.Email, Password: user.Password, Role: user.Role})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in CreateAccount GRPC API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (api *api) getQuestionHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "id is required ",
		})
		c.Error(fmt.Errorf("id is required")).SetMeta(http.StatusExpectationFailed)
		return
	}
	quest, err := api.brokerapp.GetQuestionById(c.Request.Context(), id)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in GetQuestionById  API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		//c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, quest)
}
func (api *api) getAllQuestionsHandler(c *gin.Context) {

	quests, err := api.brokerapp.GetAllQuestions(c.Request.Context())
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in GetQuestionById  API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, quests)
}
func (api *api) getAnswerHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "id is required ",
		})
		// c.Error(fmt.Errorf("id is required"))
		c.Error(fmt.Errorf("id is required")).SetMeta(http.StatusExpectationFailed)
		return
	}
	answer, err := api.brokerapp.GetAnswerByQuestionId(c.Request.Context(), id)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in GetAnswerByQuestionId  API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, answer)
}
func (api *api) getAllAnswersHandler(c *gin.Context) {
	answers, err := api.brokerapp.GetAllAnswers(c.Request.Context())
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in GetAllAnswers  API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, answers)
}
func (api *api) getAllQuestionTemplates(c *gin.Context) {
	quests, err := api.brokerapp.HandleGetAllTemplates(c.Request.Context())
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in HandleGetAllTemplates  API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, quests)
}
func (api *api) getallLanguages(c *gin.Context) {
	languages, err := api.brokerapp.GetAllLanguages(c.Request.Context())
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in GetAllLanguages API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, languages)
}
func (api *api) qcQuestion(c *gin.Context) {
	var qcPayload brokerapp.Question
	if err := c.Bind(&qcPayload); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		// c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	res, err := api.brokerapp.HandleQCQuestion(c.Request.Context(), qcPayload)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while doing question QC:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		// c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (api *api) createSnippet(c *gin.Context) {
	var snippet brokerapp.CodeSnippet
	if err := c.Bind(&snippet); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.brokerapp.HandleCreateSnippet(c.Request.Context(), snippet)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in saving snippet:", map[string]interface{}{
			"error": err.Error(),
		})
		//c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) getSnippetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "id is required ",
		})
		c.Error(fmt.Errorf("id is required")).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.brokerapp.HandleGetSnippetById(c.Request.Context(), id)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error getting snippet:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) getAllSnippetsByUser(c *gin.Context) {

}
func (api *api) formatCode(c *gin.Context) {
	var payload brokerapp.FormatterRequest
	if err := c.Bind(&payload); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.brokerapp.FormatCode(c.Request.Context(), payload)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in formatting code:", map[string]interface{}{
			"error": err.Error(),
		})
		//c.JSON(http.StatusInternalServerError, err.Error())
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) loadDbQuestion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "id is required ",
		})
		c.Error(fmt.Errorf("id is required")).SetMeta(http.StatusExpectationFailed)
		// c.JSON(http.StatusExpectationFailed, "id is required")
		return
	}
	res, err := api.brokerapp.LoadDBQuest(c.Request.Context(), id)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in loading db question:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (api *api) completeChallenge(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "id is required ",
		})
		c.Error(fmt.Errorf("id is required")).SetMeta(http.StatusExpectationFailed)
		// c.JSON(http.StatusExpectationFailed, "id is required")
		return
	}
	res, err := api.admincli.CompleteChallenge(c.Request.Context(), &admClient.CompleteChallengeRequest{ChallengeId: id})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in CompleteChallenge RPC:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) completeQuestion(c *gin.Context) {
	var payload brokerapp.CompleteChallengeRequest
	if err := c.Bind(&payload); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.admincli.CompleteQuestion(c.Request.Context(), &admClient.CompleteQuestionRequest{QuestionId: payload.QuestionId, ChallengeId: payload.ChallengeId})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in CompleteQuestion RPC:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) updateUserMetrics(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "id is required ",
		})
		c.Error(fmt.Errorf("id is required")).SetMeta(http.StatusExpectationFailed)
		// c.JSON(http.StatusExpectationFailed, "id is required")
		return
	}
	res, err := api.brokerapp.LoadDBQuest(c.Request.Context(), id)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in loading db question:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) updateUserPerformance(c *gin.Context) {
	var payload brokerapp.GlobalUserPerformance
	if err := c.Bind(&payload); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusExpectationFailed)
		return
	}
	res, err := api.admincli.UpdateUserPerformance(c.Request.Context(), &admClient.UpdateUserPerformanceRequest{})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in loading db question:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (api *api) addSubmissionStats(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": "id is required ",
		})
		c.Error(fmt.Errorf("id is required")).SetMeta(http.StatusExpectationFailed)
		// c.JSON(http.StatusExpectationFailed, "id is required")
		return
	}
	res, err := api.admincli.AddSubmissionStats(c.Request.Context(), &admClient.AddSubmissionStatsRequest{
		SubmissionStats: &admClient.SubmissionStats{},
	})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in loading db question:", map[string]interface{}{
			"error": err.Error(),
		})
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, res)
}
