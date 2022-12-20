package models

type ShippingLine struct {
	ID                            int         `json:"id"`
	CarrierIdentifier             interface{} `json:"carrier_identifier"`
	Code                          string      `json:"code"`
	DeliveryCategory              interface{} `json:"delivery_category"`
	DiscountedPrice               string      `json:"discounted_price"`
	Phone                         interface{} `json:"phone"`
	Price                         string      `json:"price"`
	RequestedFulfillmentServiceID interface{} `json:"requested_fulfillment_service_id"`
	Source                        string      `json:"source"`
	Title                         string      `json:"title"`
	TaxLines                      []Tax       `json:"tax_lines"`
	DiscountAllocations           []Discount  `json:"discount_allocations"`
}
