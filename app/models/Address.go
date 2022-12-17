package models

type Address struct {
	ID           int         `json:"id"`
	FirstName    string      `json:"first_name"`
	Address1     string      `json:"address1"`
	Phone        string      `json:"phone"`
	City         string      `json:"city"`
	Zip          string      `json:"zip"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	LastName     string      `json:"last_name"`
	Address2     string      `json:"address2"`
	Company      interface{} `json:"company"`
	Latitude     float64     `json:"latitude"`
	Longitude    float64     `json:"longitude"`
	Name         string      `json:"name"`
	CountryCode  string      `json:"country_code"`
	ProvinceCode string      `json:"province_code"`
	Default      bool        `json:"default"`
}