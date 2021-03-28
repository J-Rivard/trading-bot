package botmsgpipeline

import (
	"fmt"
	"strings"
)

func (b *BotPipeline) priceCheck() {
	for msg := range b.pcChan {
		parsed := msg.Content[1:len(msg.Content)]
		tokenized := strings.Split(parsed, " ")

		if len(tokenized) < 2 {
			b.botClient.SendMessage(msg.ChannelID, "Ticker not provided")
			continue
		}

		ticker := tokenized[1]

		stock, err := b.stockAPI.GetStockData(ticker)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			continue
		}

		b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("Price of $%s is currently %.2f", strings.ToUpper(ticker), stock.Current))
	}
}
