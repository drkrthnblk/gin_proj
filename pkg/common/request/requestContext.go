package request

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

type Context struct {
	UserID int
	Log *logrus.Entry
	clientContext context.Context
	useClientContext bool
}

func GetRequestContext(ctx *gin.Context) *Context {
	logger := ctx.Value("logger")
	return &Context{
		Log: logger.(*logrus.Entry),
		// clientContext: ctx.Request.Context(),
		useClientContext: true,
	}
}