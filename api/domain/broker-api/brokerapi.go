package brokerapi

import (
	"net/http"

	authpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto"
	execpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto"
	brokerapp "github.com/bentenison/microservice/app/domain/broker-app"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	brokerapp *brokerapp.App
	logger    *logger.CustomLogger
	authcli   authpb.AuthServiceClient
	execcli   execpb.ExecutorServiceClient
}

func newAPI(brokerApp *brokerapp.App, logger *logger.CustomLogger) *api {
	return &api{
		brokerapp: brokerApp,
		logger:    logger,
	}
}

func (api *api) newSubmissionHandler(c *gin.Context) {
	var submissionPayload SubmissionPayload
	if err := c.Bind(&submissionPayload); err != nil {
		api.logger.Errorc(c.Request.Context(), "error while binding the data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	submission := toAppSubmission(submissionPayload)
	template, err := api.brokerapp.HandleSubmisson(c.Request.Context(), submission)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	submission := toAppSubmission(submissionPayload)
	template, err := api.brokerapp.HandleCodeRun(c.Request.Context(), submission)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	// submission := toAppSubmission(submissionPayload)
	token, err := api.brokerapp.Authenticate(c.Request.Context(), cred)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting token:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}

	res, err := api.brokerapp.Authorize(c.Request.Context(), tkn.Token)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while authorizing user:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	// submission := toAppSubmission(submissionPayload)
	template, err := api.brokerapp.CreateUser(c.Request.Context(), up)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	res, err := api.authcli.Authenticate(c.Request.Context(), &authpb.LoginRequest{Username: cred.Username, Password: cred.Password})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in Authenticate GRPC API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	res, err := api.authcli.Authorize(c.Request.Context(), &authpb.AuthorizeRequest{Token: tkn.Token})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in Authorize GRPC API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	res, err := api.authcli.CreateAccount(c.Request.Context(), &authpb.CreateAccountRequest{Username: user.Username, Email: user.Email, Password: user.Password, Role: user.Role})
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in CreateAccount GRPC API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, "id is required")
		return
	}
	quest, err := api.brokerapp.GetQuestionById(c.Request.Context(), id)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in GetQuestionById  API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusExpectationFailed, "id is required")
		return
	}
	answer, err := api.brokerapp.GetAnswerByQuestionId(c.Request.Context(), id)
	if err != nil {
		api.logger.Errorc(c.Request.Context(), "error in GetAnswerByQuestionId  API:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, answers)
}
