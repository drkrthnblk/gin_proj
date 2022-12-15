package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type GenericApiError struct {
	code ErrorCode
	message string
	statusCode int
	data *[]gin.H
}

func NewGenericApiError(code ErrorCode, messae string, statusCode int, data *gin.H) error {
	return &GenericApiError{
		code: code,
		message: message,
		statusCode: statusCode,
		data: data,
	}
}

func RaiseError(ctx *gin.Context, err error) bool {
	errCode := CodeUnknown
	errMSg := string(SomethingWentWrong)
	statusCode := http.StatusInternalServerError
	data := []gin.H{}
}