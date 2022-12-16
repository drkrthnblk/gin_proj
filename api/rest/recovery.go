package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin_proj/pkg/common/response"
)

func customisedHandleRecovvery(ctx *gin.Context, err interface{}) {
	if err != nil {
		response.RaiseError(ctx, response.NewGenericApiError(response.CodePanic, "Code Panic", http.StatusInternalServerError, nil))
	}
}