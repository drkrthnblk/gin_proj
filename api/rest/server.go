package rest

import (
	"gin_proj/api/rest/endpoints"

	"github.com/gin-gonic/gin"
)

func BuildServer() *gin.Engine {
	server := gin.Default()
	server.Use(gin.CustomRecovery(customisedHandleRecovvery))
	orderGroup := server.Group("/apis")
	endpoints.OrderRoutes(orderGroup)
	return server
}
