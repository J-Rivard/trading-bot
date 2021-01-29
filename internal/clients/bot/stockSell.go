package bot

import (
	"fmt"

	"github.com/J-Rivard/trading-bot/internal/models"
)

func (b *Bot) SellStock(userID, ticker string, quantity float64) (*models.User, error) {

	user, err := b.Database.GetUser(userID)
	if err != nil {
		return nil, err
	}

	currentQuantity, ok := user.StockData[ticker]
	if !ok {
		return nil, fmt.Errorf("You don't own %s", ticker)
	}

	if currentQuantity < quantity {
		return nil, fmt.Errorf("You only own %.2f %s", currentQuantity, ticker)
	}

	stock, err := b.stockAPI.GetStockData(ticker)
	if err != nil {
		return nil, err
	}

	totalSell := stock.Current * quantity

	user.LiquidValue += totalSell
	user.StockData[ticker] -= quantity

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
