package botmsgpipeline

import (
	"fmt"
	"strings"
)

func (b *BotPipeline) buyShares() {
	for msg := range b.buySharesChan {
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

		stock, err := b.stockAPI.GetStockData(ticker)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		totalCost := stock.Current * quantityFloat

		if totalCost > user.LiquidValue {
			b.botClient.SendMessage(msg.ChannelID, "Not enough liquidity")
			continue
		}

		user.LiquidValue -= totalCost
		user.StockData[ticker] += quantityFloat

		err = b.db.UpdateUser(user)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("Remaining balance: %.2f", user.LiquidValue))
	}
}

func (b *BotPipeline) buyMoney() {
	for msg := range b.buyMoneyChan {
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

		if quantityFloat > user.LiquidValue {
			b.botClient.SendMessage(msg.ChannelID, "Not enough liquidity")
			continue
		}

		stock, err := b.stockAPI.GetStockData(ticker)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		quantityToBuy := quantityFloat / stock.Current

		user.LiquidValue -= quantityFloat
		user.StockData[ticker] += quantityToBuy

		err = b.db.UpdateUser(user)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("Purchased %.2f shares\nRemaining balance: %.2f", quantityToBuy, user.LiquidValue))
	}
}
