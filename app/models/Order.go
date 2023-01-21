package models

import (
	"gin_proj/pkg/common/utils"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Order struct {
	// ID                     int           `json:"id"`
	gorm.Model
	AdminGraphqlAPIID      string         `json:"admin_graphql_api_id"`
	AppID                  *string        `json:"app_id,omitempty"`
	BrowserIP              string         `json:"browser_ip"`
	BuyerAcceptsMarketing  bool           `json:"buyer_accepts_marketing"`
	CancelReason           *string        `json:"cancel_reason,omitempty"`
	CancelledAt            *string        `json:"cancelled_at,omitempty"`
	CartToken              string         `json:"cart_token"`
	CheckoutID             int            `json:"checkout_id"`
	CheckoutToken          string         `json:"checkout_token"`
	ClosedAt               *string        `json:"closed_at,omitempty"`
	Confirmed              bool           `json:"confirmed"`
	ContactEmail           string         `json:"contact_email"`
	CreatedAt              string         `json:"created_at"`
	Currency               string         `json:"currency"`
	CurrentSubtotalPrice   string         `json:"current_subtotal_price"`
	CurrentTotalDiscounts  string         `json:"current_total_discounts"`
	CurrentTotalDutiesSet  datatypes.JSON `json:"current_total_duties_set"`
	CurrentTotalPrice      string         `json:"current_total_price"`
	CurrentTotalTax        string         `json:"current_total_tax"`
	CustomerLocale         datatypes.JSON `json:"customer_locale"`
	DeviceID               *string        `json:"device_id,omitempty"`
	DiscountCodes          []string       `json:"discount_codes" gorm:"type:jsonb"`
	Email                  string         `json:"email"`
	FinancialStatus        string         `json:"financial_status"`
	FulfillmentStatus      *string        `json:"fulfillment_status,omitempty"`
	Gateway                string         `json:"gateway"`
	LandingSite            string         `json:"landing_site"`
	LandingSiteRef         string         `json:"landing_site_ref"`
	LocationID             *string        `json:"location_id,omitempty"`
	Name                   string         `json:"name"`
	Note                   *string        `json:"note,omitempty"`
	Number                 int            `json:"number"`
	OrderNumber            int            `json:"order_number"`
	OrderStatusURL         string         `json:"order_status_url"`
	OriginalTotalDutiesSet datatypes.JSON `json:"original_total_duties_set"`
	PaymentGatewayNames    []string       `json:"payment_gateway_names" gorm:"type:jsonb"`
	Phone                  string         `json:"phone"`
	PresentmentCurrency    string         `json:"presentment_currency"`
	ProcessedAt            string         `json:"processed_at"`
	ProcessingMethod       string         `json:"processing_method"`
	Reference              string         `json:"reference"`
	ReferringSite          string         `json:"referring_site"`
	SourceIdentifier       string         `json:"source_identifier"`
	SourceName             string         `json:"source_name"`
	SourceURL              string         `json:"source_url"`
	SubtotalPrice          string         `json:"subtotal_price"`
	Tags                   string         `json:"tags"`
	TaxLines               []Tax          `json:"tax_lines" gorm:"type:jsonb"`
	TaxesIncluded          bool           `json:"taxes_included"`
	Test                   bool           `json:"test"`
	Token                  string         `json:"token"`
	TotalDiscounts         string         `json:"total_discounts"`
	TotalLineItemsPrice    string         `json:"total_line_items_price"`
	TotalPrice             string         `json:"total_price"`
	TotalPriceUsd          string         `json:"total_price_usd"`
	TotalTax               string         `json:"total_tax"`
	TotalTipReceived       string         `json:"total_tip_received"`
	TotalWeight            int            `json:"total_weight"`
	UpdatedAt              string         `json:"updated_at"`
	UserID                 string         `json:"user_id,omitempty"`
	BillingAddress         Address        `json:"billing_address" gorm:"type:jsonb"`
	Customer               Customer       `json:"customer" gorm:"type:jsonb"`
	DiscountApplications   []Discount     `json:"discount_applications" gorm:"type:jsonb"`
	Fulfillments           []Fulfillment  `json:"fulfillments" gorm:"type:jsonb"`
	LineItems              []LineItem     `json:"line_items" gorm:"type:jsonb"`
	PaymentDetails         PaymentDetail  `json:"payment_details" gorm:"embedded"`
	Refunds                []Refund       `json:"refunds" gorm:"type:jsonb"`
	ShippingAddress        Address        `json:"shipping_address" gorm:"type:jsonb"`
	ShippingLines          []ShippingLine `json:"shipping_lines" gorm:"type:jsonb"`
}

type PaymentDetail struct {
	CreditCardBin     *string `json:"credit_card_bin,omitempty"`
	AvsResultCode     *string `json:"avs_result_code,omitempty"`
	CvvResultCode     *string `json:"cvv_result_code,omitempty"`
	CreditCardNumber  string  `json:"credit_card_number"`
	CreditCardCompany string  `json:"credit_card_company"`
}

func (t *Tax) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, t)
}

func (a *Address) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, a)
}

func (c *Customer) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, c)
}

func (d *Discount) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, d)
}

func (f *Fulfillment) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, f)
}

func (l *LineItem) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, l)
}

func (r *Refund) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, r)
}

func (s *ShippingLine) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, s)
}

func (p *PaymentDetail) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, p)
}
