package bot

import (
	"fmt"
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
		if len(tokenized) < 3 {
			return
		}

		command := tokenized[0]
		ticker := tokenized[1]
		quantity := tokenized[2]

		switch strings.ToUpper(command) {
		case Buy:
			fmt.Println(command, ticker, quantity)

		case Sell:
			fmt.Println(command, ticker, quantity)

		}
	}
}
