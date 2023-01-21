package orders

import (
	"gin_proj/app/services"
	"gin_proj/pkg/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	err := services.MigrateTables(c)
	if err != nil {
		response.RaiseError(c, err)
	}
	response.Success(c, http.StatusOK, "Migration done successfully")
}

func GetOrders(c *gin.Context) {
	orders, err := services.GetAllOrders(c)
	if err != nil {
		response.RaiseError(c, err)
	}
	response.Success(c, http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	err := services.CreateOrder(c)
	if err != nil {
		response.RaiseError(c, err)
	}
	response.Success(c, http.StatusOK, "Order creation successful")
}
