package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) SendMessage(id, msg string) (*discordgo.Message, error) {
	return b.Client.ChannelMessageSend(id, msg)
}
