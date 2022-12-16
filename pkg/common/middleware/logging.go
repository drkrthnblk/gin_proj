package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/siruspen/logrus"
	
	"gin_proj/pkg/common/request"
)

func SetIdsInCtx() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.GetHeader(string(request.HeaderRequestId))
		if requestId == "" {
			requestId = uuid.NewString()
		} else {
			_, err := uuid.Parse(requestId)
			if err != nil {
				requestId = uuid.NewString()
			}
		}

		method := ctx.Request.Method
		path := ctx.Request.URL.Path

		logger1 := logrus.WithFields(logrus.Fields{
				"request_id": requestId,
				"request_path": path,
				"request_method": method,
		})

		ctx.Set("logger", logger1)
		ctx.Next()
	}
}