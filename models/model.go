package models

type SalesSummary struct {
	StoreCode    int32  `json:"store_code"`
	Store        string `json:"store"`
	DivisionCode int32  `json:"division_code"`
	Division     string `json:"division"`
	TotalPrice   int32  `json:"total_price"`
}
