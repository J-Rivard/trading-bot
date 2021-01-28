package bot

import (
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
			if len(tokenized) < 3 {
				return
			}
			ticker := tokenized[1]
			quantity := tokenized[2]

			quantityFloat, err := strconv.ParseFloat(quantity, 64)
			if err != nil {
				fmt.Println(err)
				return
			}

			b.BuyStock(msg.Author.ID, ticker, quantityFloat)

			// stock, err := b.stockAPI.GetStockData(ticker)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }

			// quantityFloat, err := strconv.ParseFloat(quantity, 64)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }

			// b.Client.ChannelMessageSend(stonksChannelID, fmt.Sprintf("Confirmed, total cost: $%.2f", stock.Current*quantityFloat))

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
				fmt.Println(err)
				return
			}

			b.Client.ChannelMessageSend(stonksChannelID, fmt.Sprintf("Welcome to the club %s, here 100k for some tendies", msg.Author.Username))
		}
	}
}
