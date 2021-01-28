package stockapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/J-Rivard/trading-bot/internal/models"
)

const quoteEndpoint = "/api/v1/quote?"

func (s *StockAPI) GetStockData(ticker string) (*models.Stock, error) {

	endpoint := host + quoteEndpoint + fmt.Sprintf("symbol=%s&token=%s", strings.ToUpper(ticker), s.token)

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}

	var stock models.Stock
	json.NewDecoder(resp.Body).Decode(&stock)

	if stock.Current == 0 {
		return nil, errors.New("Unable to find ticker")
	}

	return &stock, nil
}
