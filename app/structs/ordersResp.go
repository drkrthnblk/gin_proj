package structs

import (
	"gin_proj/app/models"
)

type OrdersResp struct {
	Orders []models.Order `json:"orders"`
}