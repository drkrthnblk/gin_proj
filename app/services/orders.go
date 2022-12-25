package services

import(
	"fmt"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"gin_proj/app/structs"
	// "gin_proj/pkg/common/utils"
	"gin_proj/pkg/dbConns"
)

func GetAllOrders(c *gin.Context) (*structs.OrdersResp, error) {
	var orders structs.OrdersResp
	// body, err := utils.ReadRequestBody(c)
	// if err != nil {
	// 	return nil, err
	// }

	// err = json.Unmarshal(*body, &orders)
	// if err != nil {
	// 	return nil, err
	// }

	r := dbConns.NewDbConnSelectorFactory()
	fmt.Printf("%s\n", r)

	return &orders, nil
}