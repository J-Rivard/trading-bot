package bot

import (
	"github.com/bwmarrin/discordgo"
)

// This function will be called every time a new
// message is created on any channel that the authenticated bot has access to.
func (b *Bot) messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	// Handle commands from all channels
	b.outbound <- msg
}
