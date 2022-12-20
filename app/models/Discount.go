package models

type Discount struct {
	Code             string `json:"code"`
	Amount           string `json:"amount"`
	Type             string `json:"type"`
	TargetType       string `json:"target_type"`
	Value            string `json:"value"`
	ValueType        string `json:"value_type"`
	AllocationMethod string `json:"allocation_method"`
	TargetSelection  string `json:"target_selection"`
}
