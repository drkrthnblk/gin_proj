package models

type Customer struct {
	ID                        int           `json:"id"`
	Email                     string        `json:"email"`
	AcceptsMarketing          bool          `json:"accepts_marketing"`
	CreatedAt                 string        `json:"created_at"`
	UpdatedAt                 string        `json:"updated_at"`
	FirstName                 string        `json:"first_name"`
	LastName                  string        `json:"last_name"`
	OrdersCount               int           `json:"orders_count"`
	State                     string        `json:"state"`
	TotalSpent                string        `json:"total_spent"`
	LastOrderID               int           `json:"last_order_id"`
	Note                      interface{}   `json:"note"`
	VerifiedEmail             bool          `json:"verified_email"`
	MultipassIdentifier       interface{}   `json:"multipass_identifier"`
	TaxExempt                 bool          `json:"tax_exempt"`
	Tags                      string        `json:"tags"`
	LastOrderName             string        `json:"last_order_name"`
	Currency                  string        `json:"currency"`
	Phone                     string        `json:"phone"`
	AcceptsMarketingUpdatedAt string        `json:"accepts_marketing_updated_at"`
	MarketingOptInLevel       interface{}   `json:"marketing_opt_in_level"`
	TaxExemptions             []interface{} `json:"tax_exemptions"`
	AdminGraphqlAPIID         string        `json:"admin_graphql_api_id"`
	DefaultAddress            Address       `json:"default_address"`
}
