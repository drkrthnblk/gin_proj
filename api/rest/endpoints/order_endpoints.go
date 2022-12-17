package endpoints

import (
	"github.com/gin-gonic/gin"
	"gin_proj/pkg/common/middleware"
	"gin_proj/app/controllers/orders"
)

func OrderRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	orderGroup := routerGroup.Group("/orderapis")
	orderGroup.GET("/orders", middleware.SetIdsInCtx(), orders.GetOrders)
	return orderGroup
}