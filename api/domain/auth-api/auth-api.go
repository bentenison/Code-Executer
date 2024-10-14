package authapi

import (
	"context"
	"net/http"
	"time"

	"github.com/bentenison/microservice/api/domain/auth-api/grpc/proto"
	"github.com/bentenison/microservice/app/domain/authapp"
	"github.com/bentenison/microservice/app/sdk/apperrors"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	authapp *authapp.App
	log     *logger.CustomLogger
	proto.UnimplementedAuthServiceServer
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

func (s *api) loginHandler(c *gin.Context) {
	var credentials credentials
	if err := c.Bind(&credentials); err != nil {
		s.log.Errorc(c.Request.Context(), "error while binding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, apperrors.NewError(err))
	}

}
func (s *api) authorizeHandler(c *gin.Context) {

}

func (s *api) Authenticate(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	return nil, nil
}
func (s *api) CreateUser(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	return nil, nil
}
func (s *api) Authorize(ctx context.Context, req *proto.AuthorizeRequest) (*proto.AuthorizeResponse, error) {
	return nil, nil
}
