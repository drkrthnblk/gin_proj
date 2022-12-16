package endpoints

import (
	"gin_proj/pkg/common/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	// orderGroup := routerGroup.Group("/order")
	routerGroup.GET("/order", middleware.SetIdsInCtx())
	return routerGroup
}