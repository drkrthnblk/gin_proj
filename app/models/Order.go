package models

// import (
// 	"gorm.io/gorm"
// )

type Order struct {
	ID                     int           `json:"id"`
	AdminGraphqlAPIID      string        `json:"admin_graphql_api_id"`
	AppID                  interface{}   `json:"app_id"`
	BrowserIP              string        `json:"browser_ip"`
	BuyerAcceptsMarketing  bool          `json:"buyer_accepts_marketing"`
	CancelReason           interface{}   `json:"cancel_reason"`
	CancelledAt            interface{}   `json:"cancelled_at"`
	CartToken              string        `json:"cart_token"`
	CheckoutID             int           `json:"checkout_id"`
	CheckoutToken          string        `json:"checkout_token"`
	ClosedAt               interface{}   `json:"closed_at"`
	Confirmed              bool          `json:"confirmed"`
	ContactEmail           string        `json:"contact_email"`
	CreatedAt              string        `json:"created_at"`
	Currency               string        `json:"currency"`
	CurrentSubtotalPrice   string        `json:"current_subtotal_price"`
	CurrentTotalDiscounts  string        `json:"current_total_discounts"`
	CurrentTotalDutiesSet  interface{}   `json:"current_total_duties_set"`
	CurrentTotalPrice      string        `json:"current_total_price"`
	CurrentTotalTax        string        `json:"current_total_tax"`
	CustomerLocale         interface{}   `json:"customer_locale"`
	DeviceID               interface{}   `json:"device_id"`
	DiscountCodes          []interface{} `json:"discount_codes"`
	Email                  string        `json:"email"`
	FinancialStatus        string        `json:"financial_status"`
	FulfillmentStatus      interface{}   `json:"fulfillment_status"`
	Gateway                string        `json:"gateway"`
	LandingSite            string        `json:"landing_site"`
	LandingSiteRef         string        `json:"landing_site_ref"`
	LocationID             interface{}   `json:"location_id"`
	Name                   string        `json:"name"`
	Note                   interface{}   `json:"note"`
	Number                 int           `json:"number"`
	OrderNumber            int           `json:"order_number"`
	OrderStatusURL         string        `json:"order_status_url"`
	OriginalTotalDutiesSet interface{}   `json:"original_total_duties_set"`
	PaymentGatewayNames    []string      `json:"payment_gateway_names"`
	Phone                  string        `json:"phone"`
	PresentmentCurrency    string        `json:"presentment_currency"`
	ProcessedAt            string        `json:"processed_at"`
	ProcessingMethod       string        `json:"processing_method"`
	Reference              string        `json:"reference"`
	ReferringSite          string        `json:"referring_site"`
	SourceIdentifier       string        `json:"source_identifier"`
	SourceName             string        `json:"source_name"`
	SourceURL              interface{}   `json:"source_url"`
	SubtotalPrice          string        `json:"subtotal_price"`
	Tags                   string        `json:"tags"`
	TaxLines               []Tax         `json:"tax_lines" gorm:"type:jsonb"`
	TaxesIncluded          bool          `json:"taxes_included"`
	Test                   bool          `json:"test"`
	Token                  string        `json:"token"`
	TotalDiscounts         string        `json:"total_discounts"`
	TotalLineItemsPrice    string        `json:"total_line_items_price"`
	TotalPrice             string        `json:"total_price"`
	TotalPriceUsd          string        `json:"total_price_usd"`
	TotalTax               string        `json:"total_tax"`
	TotalTipReceived       string        `json:"total_tip_received"`
	TotalWeight            int           `json:"total_weight"`
	UpdatedAt              string        `json:"updated_at"`
	UserID                 interface{}   `json:"user_id"`
	BillingAddress         Address       `json:"billing_address" gorm:"type:jsonb"`
	Customer               Customer      `json:"customer"`
	DiscountApplications   []Discount    `json:"discount_applications" gorm:"type:jsonb"`
	Fulfillments           []Fulfillment `json:"fulfillments" gorm:"type:jsonb"`
	LineItems              []LineItem    `json:"line_items" gorm:"type:jsonb"`
	PaymentDetails         struct {
		CreditCardBin     interface{} `json:"credit_card_bin"`
		AvsResultCode     interface{} `json:"avs_result_code"`
		CvvResultCode     interface{} `json:"cvv_result_code"`
		CreditCardNumber  string      `json:"credit_card_number"`
		CreditCardCompany string      `json:"credit_card_company"`
	} `json:"payment_details"`
	Refunds         []Refund       `json:"refunds" gorm:"type:jsonb"`
	ShippingAddress Address        `json:"shipping_address" gorm:"type:jsonb"`
	ShippingLines   []ShippingLine `json:"shipping_lines" gorm:"type:jsonb"`
}
