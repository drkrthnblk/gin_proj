package services

import (
	"encoding/json"
	"fmt"
	"gin_proj/app/models"
	postgresrepo "gin_proj/app/repository/postgresRepo"
	"gin_proj/app/structs"

	"github.com/gin-gonic/gin"

	"gin_proj/pkg/common/request"
	"gin_proj/pkg/common/utils"
	dbconns "gin_proj/pkg/dbConns"
	"gin_proj/pkg/dbConns/postgresConn"
)

func MigrateTables(c *gin.Context) error {
	var order models.Order
	rCtx := request.GetRequestContext(c)
	orderRepo := postgresrepo.OrderRepo{}
	err := orderRepo.Migrate(rCtx, order)
	if err != nil {
		return err
	}
	return nil
}
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
	fmt.Printf("%s\n", r.(*postgresConn.PostgresClient))

	return &orders, nil
}

func CreateOrder(c *gin.Context) error {
	var order models.Order
	body, err := utils.ReadRequestBody(c)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*body, &order)
	if err != nil {
		return err
	}

	rCtx := request.GetRequestContext(c)
	orderRepo := postgresrepo.OrderRepo{}
	err = orderRepo.Create(rCtx, order)
	if err != nil {
		return err
	}
	return nil
}
