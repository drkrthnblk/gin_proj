package models

import (
	"gorm.io/gorm"
)

type Fulfillment struct {
	// ID                int    `json:"id"`
	gorm.Model
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	CreatedAt         string `json:"created_at"`
	LocationID        int    `json:"location_id"`
	Name              string `json:"name"`
	OrderID           int    `json:"order_id"`
	Receipt           struct {
		Testcase      bool   `json:"testcase"`
		Authorization string `json:"authorization"`
	} `json:"receipt" gorm:"embedded"`
	Service         string     `json:"service"`
	ShipmentStatus  string     `json:"shipment_status"`
	Status          string     `json:"status"`
	TrackingCompany string     `json:"tracking_company"`
	TrackingNumber  string     `json:"tracking_number"`
	TrackingNumbers []string   `json:"tracking_numbers" gorm:"type:jsonb"`
	TrackingURL     string     `json:"tracking_url"`
	TrackingUrls    []string   `json:"tracking_urls" gorm:"type:jsonb"`
	UpdatedAt       string     `json:"updated_at"`
	LineItems       []LineItem `json:"line_items" gorm:"type:jsonb"`
}
