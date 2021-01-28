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

func (s *StockAPI) CalculateUserValue(user *models.User) (float64, error) {
	var totalAssetValue float64

	for k, v := range user.StockData {
		stock, err := s.GetStockData(k)
		if err != nil {
			return 0, err
		}

		totalAssetValue += (stock.Current * v)
	}

	return totalAssetValue, nil
}
