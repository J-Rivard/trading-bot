package db

import (
	"encoding/json"

	"github.com/J-Rivard/trading-bot/internal/models"
)

func (d *DB) SubscribeUser(user *models.User) error {
	insertString := `INSERT INTO trading_bot.users
	 (id, liquid_value, asset_value, stock_data)
	 VALUES ($1, $2, $3, $4)`

	stockData, err := json.Marshal(user.StockData)
	if err != nil {
		return err
	}

	_, err = d.Client.Exec(insertString, user.ID, user.LiquidValue, user.AssetValue, stockData)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetUser(id string) (*models.User, error) {
	query := `SELECT * FROM trading_bot.users WHERE id=$1`

	row := d.Client.QueryRow(query, id)

	var stockData []byte
	var user models.User

	err := row.Scan(&user.ID, &user.LiquidValue, &user.AssetValue, &stockData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(stockData, &user.StockData)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
