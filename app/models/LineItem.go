package models

type LineItem struct {
	ID                  int         `json:"id"`
	AdminGraphqlAPIID   string      `json:"admin_graphql_api_id"`
	FulfillableQuantity int         `json:"fulfillable_quantity"`
	FulfillmentService  string      `json:"fulfillment_service"`
	FulfillmentStatus   interface{} `json:"fulfillment_status"`
	GiftCard            bool        `json:"gift_card"`
	Grams               int         `json:"grams"`
	Name                string      `json:"name"`
	Price               string      `json:"price"`
	ProductExists       bool        `json:"product_exists"`
	ProductID           int         `json:"product_id"`
	Properties          []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
	Quantity                   int         `json:"quantity"`
	RequiresShipping           bool        `json:"requires_shipping"`
	Sku                        string      `json:"sku"`
	Taxable                    bool        `json:"taxable"`
	Title                      string      `json:"title"`
	TotalDiscount              string      `json:"total_discount"`
	VariantID                  int         `json:"variant_id"`
	VariantInventoryManagement string      `json:"variant_inventory_management"`
	VariantTitle               string      `json:"variant_title"`
	Vendor                     interface{} `json:"vendor"`
	TaxLines                   []struct {
		Price string  `json:"price"`
		Rate  float64 `json:"rate"`
		Title string  `json:"title"`
	} `json:"tax_lines"`
	Duties              []interface{} `json:"duties"`
	DiscountAllocations []struct {
		Amount                   string `json:"amount"`
		DiscountApplicationIndex int    `json:"discount_application_index"`
	} `json:"discount_allocations"`
}