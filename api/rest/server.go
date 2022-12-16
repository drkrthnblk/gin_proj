package rest

import(
	"github.com/gin-gonic/gin"
	"gin_proj/api/rest/endpoints"
)

func BuildServer() *gin.Engine {
	server := gin.Default()
	server.Use(gin.CustomRecovery(customisedHandleRecovvery))
	orderGroup := server.Group("/orders")
	endpoints.OrderRoutes(orderGroup)
	return server
}