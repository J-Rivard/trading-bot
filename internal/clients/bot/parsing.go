package bot

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleCommand(msg *discordgo.MessageCreate) {
	if !(msg.ChannelID == stonksChannelID || msg.ChannelID == stonksDevChannelID) {
		return
	}

	if len(msg.Content) < 1 {
		return
	}

	if string(msg.Content[0]) == Prefix {
		parsed := msg.Content[1:len(msg.Content)]
		tokenized := strings.Split(parsed, " ")
		if len(tokenized) < 1 {
			return
		}

		command := tokenized[0]

		switch strings.ToUpper(command) {
		case Buy:
			ticker, quantityFloat, err := extractTickerQuantity(tokenized)
			if err != nil {
				b.Client.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}

			user, err := b.BuyStock(msg.Author.ID, ticker, quantityFloat)
			if err != nil {
				b.Client.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}

			b.Client.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Liquidity: $%.2f, Assets $%.2f", user.LiquidValue, user.AssetValue))

		case Sell:
			ticker, quantityFloat, err := extractTickerQuantity(tokenized)
			if err != nil {
				b.Client.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}

			user, err := b.SellStock(msg.Author.ID, ticker, quantityFloat)
			if err != nil {
				b.Client.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}

			b.Client.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Liquidity: $%.2f, Assets $%.2f", user.LiquidValue, user.AssetValue))
		case Join:
			err := b.SubscribeUser(msg.Author.ID)
			if err != nil {
				b.Client.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}

			b.Client.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Welcome to the club %s, here 100k for some tendies", msg.Author.Username))

		case Stats:
			user, err := b.Database.GetUser(msg.Author.ID)
			if err != nil {
				b.Client.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}

			assetVal, err := b.stockAPI.CalculateUserValue(user)
			if err != nil {
				b.Client.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}
			b.Database.UpdateUser(user)

			stockString := ""
			for k, v := range user.StockData {
				stockString += fmt.Sprintf("%s: %.2f\n", k, v)
			}

			b.Client.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Current valuation: %.2f\n"+
				fmt.Sprintf("Liquidity: %.2f\n", user.LiquidValue)+
				"Stocks:\n"+
				stockString,
				user.LiquidValue+assetVal))
		case Help:
			b.Client.ChannelMessageSend(msg.ChannelID, "Available commands:\n"+
				"!join\n"+
				"!buy <ticker> <quantity>\n"+
				"!sell <ticker> <quantity>\n"+
				"!stats")
		}
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
