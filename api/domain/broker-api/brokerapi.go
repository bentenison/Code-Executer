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

func (api *api) getQuestionTemplate(c *gin.Context) {
	template, err := api.brokerapp.GetTemplate(c.Request.Context())
	if err != nil {
		api.logger.Error("error while getting template data:", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, template)
}
