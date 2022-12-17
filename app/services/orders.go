package services

import(
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gin_proj/app/structs"
	"gin_proj/pkg/common/utils"
)

func GetAllOrders(c *gin.Context) (*structs.OrdersResp, error) {
	var orders structs.OrdersResp
	body, err := utils.ReadRequestBody(c)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*body, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}