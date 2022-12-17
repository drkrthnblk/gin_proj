package models

type Tax struct {
	Price string  `json:"price"`
	Rate  float64 `json:"rate"`
	Title string  `json:"title"`
}