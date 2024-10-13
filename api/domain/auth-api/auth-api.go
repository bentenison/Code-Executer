package authapi

import (
	"net/http"
	"time"

	"github.com/bentenison/microservice/app/domain/authapp"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	authapp *authapp.App
	log     *logger.CustomLogger
}

func newAPI(log *logger.CustomLogger, authapp *authapp.App) *api {
	return &api{
		authapp: authapp,
		log:     log,
	}
}

func (a *api) checkHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "running",
		"time":   time.Now(),
	})
}
