package stockapi

import (
	"net/http"

	"github.com/J-Rivard/trading-bot/internal/logging"
)

type StockAPI struct {
	token  string
	client *http.Client
}

type Parameters struct {
	Token string
}

const (
	host = "https://finnhub.io"
)

func New(params *Parameters, log *logging.Log) (*StockAPI, error) {
	return &StockAPI{
			token:  params.Token,
			client: &http.Client{},
		},
		nil
}
