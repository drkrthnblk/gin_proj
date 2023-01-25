package models

import (
	"gorm.io/gorm"
)

type ShippingLine struct {
	// ID                            int        `json:"id"`
	gorm.Model
	CarrierIdentifier             string     `json:"carrier_identifier,omitempty"`
	Code                          string     `json:"code"`
	DeliveryCategory              string     `json:"delivery_category,omitempty"`
	DiscountedPrice               string     `json:"discounted_price"`
	Phone                         string     `json:"phone,omitempty"`
	Price                         string     `json:"price"`
	RequestedFulfillmentServiceID string     `json:"requested_fulfillment_service_id,omitempty"`
	Source                        string     `json:"source"`
	Title                         string     `json:"title"`
	TaxLines                      []Tax      `json:"tax_lines" gorm:"type:jsonb"`
	DiscountAllocations           []Discount `json:"discount_allocations" gorm:"type:jsonb"`
}
