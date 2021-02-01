package botmsgpipeline

import (
	"fmt"

	"github.com/J-Rivard/trading-bot/internal/models"
)

func (b *BotPipeline) join() {
	for msg := range b.joinChan {
		user := models.NewUser(msg.Author.ID)

		err := b.db.SubscribeUser(user)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("Welcome to the club %s, here 100k for some tendies", msg.Author.Username))
	}
}

func (b *BotPipeline) stats() {
Loop:
	for msg := range b.statsChan {
		user, err := b.db.GetUser(msg.Author.ID)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		var totalAsset float64

		stockString := ""
		for k, v := range user.StockData {
			stock, err := b.stockAPI.GetStockData(k)
			if err != nil {
				b.botClient.SendMessage(msg.ChannelID, err.Error())
				continue Loop
			}

			stockString += fmt.Sprintf("%s: %.2f shares, $%.2f\n", k, v, stock.Current*v)
			totalAsset += stock.Current * v
		}

		b.botClient.SendMessage(msg.ChannelID,
			fmt.Sprintf(
				fmt.Sprintf("Current valuation: %.2f\n", user.LiquidValue+totalAsset)+
					fmt.Sprintf("Liquidity: %.2f\n", user.LiquidValue)+
					fmt.Sprintf("Asset value: %.2f\n", totalAsset)+
					fmt.Sprintf("Stocks: %s\n", stockString),
			),
		)

	}
}
