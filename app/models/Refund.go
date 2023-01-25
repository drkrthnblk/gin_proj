package models

import (
	"gorm.io/gorm"
)

type Refund struct {
	// ID                int           `json:"id"`
	gorm.Model
	AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
	CreatedAt         string        `json:"created_at"`
	Note              string        `json:"note"`
	OrderID           int           `json:"order_id"`
	ProcessedAt       string        `json:"processed_at"`
	Restock           bool          `json:"restock"`
	UserID            int           `json:"user_id"`
	OrderAdjustments  []interface{} `json:"order_adjustments" gorm:"type:jsonb"`
	Transactions      []struct {
		ID                int    `json:"id"`
		AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
		Amount            string `json:"amount"`
		Authorization     string `json:"authorization"`
		CreatedAt         string `json:"created_at"`
		Currency          string `json:"currency"`
		DeviceID          string `json:"device_id,omitempty"`
		ErrorCode         string `json:"error_code,omitempty"`
		Gateway           string `json:"gateway"`
		Kind              string `json:"kind"`
		LocationID        string `json:"location_id,omitempty"`
		Message           string `json:"message,omitempty"`
		OrderID           int    `json:"order_id"`
		ParentID          int    `json:"parent_id"`
		ProcessedAt       string `json:"processed_at"`
		Receipt           struct {
		} `json:"receipt"`
		SourceName string `json:"source_name"`
		Status     string `json:"status"`
		Test       bool   `json:"test"`
		UserID     string `json:"user_id,omitempty"`
	} `json:"transactions"`
	RefundLineItems []struct {
		ID          int      `json:"id"`
		LineItemID  int      `json:"line_item_id"`
		LocationID  int      `json:"location_id"`
		Quantity    int      `json:"quantity"`
		RestockType string   `json:"restock_type"`
		Subtotal    float64  `json:"subtotal"`
		TotalTax    float64  `json:"total_tax"`
		LineItem    LineItem `json:"line_item" gorm:"type:jsonb"`
	} `json:"refund_line_items"`
	Duties []interface{} `json:"duties" gorm:"type:jsonb"`
}
