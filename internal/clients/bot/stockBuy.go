package bot

import (
	"errors"

	"github.com/J-Rivard/trading-bot/internal/models"
)

func (b *Bot) BuyStock(userID, ticker string, quantity float64) (*models.User, error) {

	user, err := b.Database.GetUser(userID)
	if err != nil {
		return nil, err
	}

	stock, err := b.stockAPI.GetStockData(ticker)
	if err != nil {
		return nil, err
	}

	totalCost := stock.Current * quantity

	if totalCost > user.LiquidValue {
		return nil, errors.New("Not enough liquidity")
	}

	user.LiquidValue -= totalCost
	user.StockData[ticker] += quantity

	user.AssetValue, err = b.stockAPI.CalculateUserValue(user)
	if err != nil {
		return nil, err
	}

	err = b.Database.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
