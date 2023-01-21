package endpoints

import (
	"gin_proj/app/controllers/orders"
	"gin_proj/pkg/common/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	orderGroup := routerGroup.Group("/orderapis")
	orderGroup.GET("/migrate", middleware.SetIdsInCtx(), orders.Migrate)
	orderGroup.GET("/orders", middleware.SetIdsInCtx(), orders.GetOrders)
	orderGroup.POST("/order", middleware.SetIdsInCtx(), orders.CreateOrder)
	return orderGroup
}
