package stockapi

import "github.com/J-Rivard/trading-bot/internal/logging"

type StockAPI struct {
}

type Parameters struct {
	Token string
}

func New(params *Parameters, log *logging.Log) (*StockAPI, error) {
	return &StockAPI{},
		nil
}
