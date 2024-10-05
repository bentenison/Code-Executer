package mid

import (
	"github.com/gin-gonic/gin"
)

const (
	claimKey  = "claimctx"
	userIDKey = "userIdctx"
	userKey   = "userctx"
	bookKey   = "bookctx"
	trKey     = "trsnctx"
	traceKey  = "tracectx"
)

func GetTraceId(ctx *gin.Context) (any, bool) {
	return ctx.Get(traceKey)
}
