package models


type Fulfillment struct {
	ID                int    `json:"id"`
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	CreatedAt         string `json:"created_at"`
	LocationID        int    `json:"location_id"`
	Name              string `json:"name"`
	OrderID           int    `json:"order_id"`
	Receipt           struct {
		Testcase      bool   `json:"testcase"`
		Authorization string `json:"authorization"`
	} `json:"receipt"`
	Service         string        `json:"service"`
	ShipmentStatus  interface{}   `json:"shipment_status"`
	Status          string        `json:"status"`
	TrackingCompany string        `json:"tracking_company"`
	TrackingNumber  string        `json:"tracking_number"`
	TrackingNumbers []string      `json:"tracking_numbers"`
	TrackingURL     string        `json:"tracking_url"`
	TrackingUrls    []string      `json:"tracking_urls"`
	UpdatedAt       string        `json:"updated_at"`
	LineItems       []LineItem `json:"line_items"`
}