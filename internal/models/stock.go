package models

type Stock struct {
	Current float64 `json:"c"`
	High    float64 `json:"h"`
	Low     float64 `json:"l"`
	Open    float64 `json:"o"`
}
