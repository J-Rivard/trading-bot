package bot

import (
	"fmt"

	"github.com/J-Rivard/trading-bot/internal/logging"

	"github.com/bwmarrin/discordgo"
)

const (
	Prefix = "!"
	Buy    = "BUY"
	Sell   = "SELL"
)

// This function will be called every time a new
// message is created on any channel that the authenticated bot has access to.
func (b *Bot) messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	// Handle commands from all channels
	b.handleCommand(msg)
}

// This function will be called every time a
// message is updated on any channel that the authenticated bot has access to.
func (b *Bot) messageUpdate(session *discordgo.Session, msg *discordgo.MessageUpdate) {
	if msg.Author == nil || msg.Author.ID == session.State.User.ID {
		return
	}

	for _, embed := range msg.Embeds {
		for _, field := range embed.Fields {
			b.Log.LogDebug(logging.FormattedLog{
				"action":   "messageUpdate",
				"metadata": fmt.Sprintf("%s\n", field.Value),
			})
		}
	}
}
