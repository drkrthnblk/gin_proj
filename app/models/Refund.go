package models

type Refund struct {
	ID                int           `json:"id"`
	AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
	CreatedAt         string        `json:"created_at"`
	Note              string        `json:"note"`
	OrderID           int           `json:"order_id"`
	ProcessedAt       string        `json:"processed_at"`
	Restock           bool          `json:"restock"`
	UserID            int           `json:"user_id"`
	OrderAdjustments  []interface{} `json:"order_adjustments"`
	Transactions      []struct {
		ID                int         `json:"id"`
		AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
		Amount            string      `json:"amount"`
		Authorization     string      `json:"authorization"`
		CreatedAt         string      `json:"created_at"`
		Currency          string      `json:"currency"`
		DeviceID          interface{} `json:"device_id"`
		ErrorCode         interface{} `json:"error_code"`
		Gateway           string      `json:"gateway"`
		Kind              string      `json:"kind"`
		LocationID        interface{} `json:"location_id"`
		Message           interface{} `json:"message"`
		OrderID           int         `json:"order_id"`
		ParentID          int         `json:"parent_id"`
		ProcessedAt       string      `json:"processed_at"`
		Receipt           struct {
		} `json:"receipt"`
		SourceName string      `json:"source_name"`
		Status     string      `json:"status"`
		Test       bool        `json:"test"`
		UserID     interface{} `json:"user_id"`
	} `json:"transactions"`
	RefundLineItems []struct {
		ID          int     `json:"id"`
		LineItemID  int     `json:"line_item_id"`
		LocationID  int     `json:"location_id"`
		Quantity    int     `json:"quantity"`
		RestockType string  `json:"restock_type"`
		Subtotal    float64 `json:"subtotal"`
		TotalTax    float64 `json:"total_tax"`
		LineItem    LineItem `json:"line_item"`
	} `json:"refund_line_items"`
	Duties []interface{} `json:"duties"`
}