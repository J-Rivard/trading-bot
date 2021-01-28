package models

type User struct {
	ID          string
	LiquidValue float64
	AssetValue  float64
	TotalValue  float64
	StockData   map[string]float64
}

func NewUser(id string) *User {
	return &User{
		ID:          id,
		LiquidValue: 100000,
		StockData:   make(map[string]float64),
	}
}
