package botmsgpipeline

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	Prefix    = "!"
	BuyShares = "BUYSHARES"
	BuyMoney  = "BUYMONEY"
	Sell      = "SELL"
	Join      = "JOIN"
	Stats     = "STATS"
	Help      = "HELP"

	stonksChannelID    = "804809948583559220"
	stonksDevChannelID = "804911890484428830"
)

func (b *BotPipeline) Start() {

}

func (b *BotPipeline) validate() {
	for msg := range b.validateChan {
		if !(msg.ChannelID == stonksChannelID || msg.ChannelID == stonksDevChannelID) {
			continue
		}

		if len(msg.Content) < 1 {
			continue
		}

		if string(msg.Content[0]) == Prefix {
			parsed := msg.Content[1:len(msg.Content)]
			tokenized := strings.Split(parsed, " ")
			if len(tokenized) < 1 {
				continue
			}
		}

		b.parseChan <- msg
	}
}

func (b *BotPipeline) parse() {
	for msg := range b.parseChan {
		parsed := msg.Content[1:len(msg.Content)]
		tokenized := strings.Split(parsed, " ")
		command := tokenized[0]

		switch strings.ToUpper(command) {
		case BuyShares:
			b.buySharesChan <- msg

		case BuyMoney:
			b.buyMoneyChan <- msg

		case Sell:
			b.sellSharesChan <- msg

		case Join:
			b.joinChan <- msg

		case Stats:
			b.statsChan <- msg

		case Help:
			b.helpChan <- msg
		}
	}
}

func (b *BotPipeline) buyShares() {
	for msg := range b.buySharesChan {
		parsed := msg.Content[1:len(msg.Content)]
		tokenized := strings.Split(parsed, " ")

		ticker, quantityFloat, err := extractTickerQuantity(tokenized)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			return
		}

		user, err := b.BuyStock(msg.Author.ID, ticker, quantityFloat)
		if err != nil {
			b.botClient.SendMessage(msg.ChannelID, err.Error())
			return
		}

		b.botClient.SendMessage(msg.ChannelID, fmt.Sprintf("Liquidity: $%.2f, Assets $%.2f", user.LiquidValue, user.AssetValue))
	}
}

func extractTickerQuantity(tokenized []string) (string, float64, error) {
	if len(tokenized) < 3 {
		return "", 0, errors.New("Not enough args")
	}
	ticker := strings.ToUpper(tokenized[1])
	quantity := tokenized[2]

	quantityFloat, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return "", 0, err
	}

	return ticker, quantityFloat, nil
}
