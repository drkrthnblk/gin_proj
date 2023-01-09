package services

import (
	"fmt"
	// "encoding/json"
	"gin_proj/app/structs"

	"github.com/gin-gonic/gin"

	// "gin_proj/pkg/common/utils"
	dbconns "gin_proj/pkg/dbConns"
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

	r, _ := dbconns.NewDbConnSelectorFactory().GetDbConn(dbconns.POSTGRESCONN)
	fmt.Printf("%s\n", r)

	return &orders, nil
}
