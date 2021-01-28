package bot

import (
	"errors"
)

func (b *Bot) BuyStock(userID, ticker string, quantity float64) error {

	user, err := b.Database.GetUser(userID)
	if err != nil {
		return err
	}

	stock, err := b.stockAPI.GetStockData(ticker)
	if err != nil {
		return err
	}

	totalCost := stock.Current * quantity

	if totalCost > user.LiquidValue {
		return errors.New("Not enough liquidity")
	}

	user.LiquidValue -= totalCost
	user.AssetValue += totalCost

	return nil
}
