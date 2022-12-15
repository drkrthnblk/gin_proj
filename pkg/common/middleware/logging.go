package middleware

import (
	"gin_proj/pkg/common/request"
	"github.com/google/uuid"
	"github.com/siruspen/logrus"
)

func SetIdsInCtx() gin.HandlerFuc {
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
		path := ctx.Request.Path

		logger := logrus.WithFields(
			logrus.Fields{
				"request_id": requestId,
				"request_path": path,
				"request_method": method,
			}
		)
		ctx.Set("logger", logger)
		ctx.Next()
	}
}