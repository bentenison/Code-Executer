package examapi

import (
	"github.com/bentenison/microservice/app/domain/examapp"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	examApp *examapp.App
	logger  *logger.CustomLogger
}

func newAPI(examapp *examapp.App, logger *logger.CustomLogger) *api {
	return &api{
		examApp: examapp,
		logger:  logger,
	}
}

func (a *api) Handler(c *gin.Context) {

}
