package bot

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleCommand(msg *discordgo.MessageCreate) {
	if msg.ChannelID != stonksChannelID {
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
				b.Client.ChannelMessageSend(stonksChannelID, err.Error())
				return
			}

			user, err := b.BuyStock(msg.Author.ID, ticker, quantityFloat)
			if err != nil {
				b.Client.ChannelMessageSend(stonksChannelID, err.Error())
				return
			}

			b.Client.ChannelMessageSend(stonksChannelID, fmt.Sprintf("Liquidity: $%.2f, Assets $%.2f", user.LiquidValue, user.AssetValue))

		case Sell:
			if len(tokenized) < 3 {
				return
			}
			ticker := tokenized[1]
			quantity := tokenized[2]
			fmt.Println(command, ticker, quantity)
		case Join:
			err := b.SubscribeUser(msg.Author.ID)
			if err != nil {
				b.Client.ChannelMessageSend(stonksChannelID, err.Error())
				return
			}

			b.Client.ChannelMessageSend(stonksChannelID, fmt.Sprintf("Welcome to the club %s, here 100k for some tendies", msg.Author.Username))

		case Stats:
			user, err := b.Database.GetUser(msg.Author.ID)
			if err != nil {
				b.Client.ChannelMessageSend(stonksChannelID, err.Error())
				return
			}

			assetVal, err := b.stockAPI.CalculateUserValue(user)
			if err != nil {
				b.Client.ChannelMessageSend(stonksChannelID, err.Error())
				return
			}

			b.Client.ChannelMessageSend(stonksChannelID, fmt.Sprintf("Current valuation: %.2f", user.LiquidValue+assetVal))
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
