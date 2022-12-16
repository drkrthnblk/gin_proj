package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type GenericApiError struct {
	code ErrorCode
	message ErrorMessages
	statusCode int
	data *[]gin.H
}

func (e *GenericApiError) Error() string {
	return string(e.message)
}

func (e *GenericApiError) Code() string {
	return string(e.code)
}

func (e *GenericApiError) StatusCode() int {
	return e.statusCode
}

func (e *GenericApiError) Data() *[]gin.H {
	return e.data
}

func NewGenericApiError(code ErrorCode, message ErrorMessages, statusCode int, data *[]gin.H) error {
	return &GenericApiError{
		code: code,
		message: message,
		statusCode: statusCode,
		data: data,
	}
}

func RaiseError(ctx *gin.Context, err error) bool {
	errCode := CodeUnknown
	errMSg := string(ErrSomethingWentWrong)
	statusCode := http.StatusInternalServerError
	data := []gin.H{}
	switch e := err.(type) {
	case *GenericApiError:
		errCode = ErrorCode(e.Code())
		errMSg = e.Error()
		statusCode = e.StatusCode()
		if e.Data() != nil {
			data = *e.Data()
		}
		switch statusCode{
		case 499:
			ctx.AbortWithStatusJSON(499, map[string]interface{}{
				"status": "fail",
				"data": nil,
				"err": string(errCode),
				"msg": errMSg,
			})
		case http.StatusInternalServerError, http.StatusNotImplemented:
			Error(
				ctx,
				statusCode,
				data,
				errCode,
				errMSg,
			)
		default:
			Fail(
				ctx,
				statusCode,
				data,
				errCode,
				errMSg,
			)
		}
		
	default:
		data = append(data, gin.H{"error": err.Error()})
		Error(
			ctx,
			statusCode,
			data,
			errCode,
			errMSg,
		)
	}
	return true
}