package response

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

type response gin.H

func Success(ctx *gin.Context, statusCode int, data interface{}) {
	logger := ctx.Value("logger")
	log := logger.(*logrus.Entry)
	if ctx.Writer.Written() {
		log.Warn("response body was already written! wil not overwrite")
		return
	}

	res := response{
		"statusCode": 200,
		"data": data,
		"code": "",
		"message": "",
	}
	ctx.JSON(statusCode, res)
}

func Fail(ctx *gin.Context, statusCode int, errors []gin.H, errCode ErrorCode, msg string) {
	logger := ctx.Value("logger")
	log := logger.(*logrus.Entry)
	if ctx.Writer.Written() {
		log.Warn("response body was already written! wil not overwrite")
		return
	}

	res := response{
		"statusCode": statusCode,
		"data": response{
			"errors": errors,
		},
		"code": statusCode,
		"message": msg,
	}
	ctx.JSON(statusCode, res)
}

func Error(ctx *gin.Context, statusCode int, errors []gin.H, errCode ErrorCode, msg string) {
	logger := ctx.Value("logger")
	log := logger.(*logrus.Entry)
	if ctx.Writer.Written() {
		log.Warn("response body was already written! wil not overwrite")
		return
	}

	res := response{
		"statusCode": statusCode,
		"data": response{
			"errors": errors,
		},
		"code": errCode,
		"message": msg,
	}
	ctx.JSON(statusCode, res)
}

func SetCookie(ctx *gin.Context, name, value, path, host string, maxAe int) {
	hostType := os.Getenv("HOST_TYPE")
	if hostType == "" {
		hostType = "local"
	}
	if hostType == "stage" {
		ctx.SetCookie(name, value, maxAe, path, "localhost", false, false)
	}
	ctx.SetCookie(name, value, maxAe, path, host, true, true)
}