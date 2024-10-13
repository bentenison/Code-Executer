package brokerapi

import (
	"net/http"

	brokerapp "github.com/bentenison/microservice/app/domain/broker-app"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	brokerapp *brokerapp.App
	logger    *logger.CustomLogger
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
