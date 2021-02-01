package botmsgpipeline

import (
	"fmt"
	"strings"
)

func (b *BotPipeline) sellShares() {
	for msg := range b.sellSharesChan {
		parsed := msg.Content[1:len(msg.Content)]
		tokenized := strings.Split(parsed, " ")

		ticker, quantityFloat, err := extractTickerQuantity(tokenized)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		user, err := b.db.GetUser(msg.Author.ID)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		currentQuantity, ok := user.StockData[ticker]
		if !ok {
			b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("You don't own %s", ticker))
			continue
		}

		if currentQuantity < quantityFloat {
			b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("You only own %.2f %s", currentQuantity, ticker))
			continue
		}

		stock, err := b.stockAPI.GetStockData(ticker)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		totalSell := stock.Current * quantityFloat

		user.LiquidValue += totalSell
		user.StockData[ticker] -= quantityFloat

		if user.StockData[ticker] == 0 {
			delete(user.StockData, ticker)
		}

		err = b.db.UpdateUser(user)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("Balance: %.2f", user.LiquidValue))
	}
}
