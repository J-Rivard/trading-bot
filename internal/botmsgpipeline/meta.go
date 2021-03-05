package botmsgpipeline

func (b *BotPipeline) help() {
	for msg := range b.helpChan {
		b.botClient.SendMessage(msg.ChannelID, "Available commands:\n"+
			"$join\n"+
			"$buy <ticker> <quantity>\n"+
			"$buymoney <ticker> <dollars>\n"+
			"$sell <ticker> <quantity>\n"+
			"$stats")
	}
}
