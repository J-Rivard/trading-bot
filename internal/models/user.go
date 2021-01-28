package models

type User struct {
	ID          string
	LiquidValue float64
	AssetValue  float64
	TotalValue  float64
	StockData   []*UserStockData
}

type UserStockData struct {
	Ticker   string
	Quantity float64
}

func NewUser(id string) *User {
	return &User{
		ID:          id,
		LiquidValue: 100000,
		StockData:   []*UserStockData{},
	}
}
