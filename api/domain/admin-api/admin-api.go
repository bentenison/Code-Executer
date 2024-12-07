package adminapi

import (
	"github.com/bentenison/microservice/app/domain/adminapp"
	"github.com/bentenison/microservice/foundation/logger"
)

type api struct {
	adminapp *adminapp.App
	logger   *logger.CustomLogger
	// authcli   authpb.AuthServiceClient
	// execcli   execpb.ExecutorServiceClient
}

func newAPI(adminApp *adminapp.App, logger *logger.CustomLogger) *api {
	return &api{
		adminapp: adminApp,
		logger:   logger,
	}
}
