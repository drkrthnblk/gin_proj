package orders

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"gin_proj/app/services"
	"gin_proj/pkg/common/response"
)

func GetOrders(c *gin.Context) {
	orders, err := services.GetAllOrders(c)
	if err != nil {
		response.RaiseError(c, err)
	}
	response.Success(c, http.StatusOK, orders)
}

