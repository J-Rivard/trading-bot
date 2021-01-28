package bot

import "github.com/J-Rivard/trading-bot/internal/models"

func (b *Bot) SubscribeUser(userID string) error {
	user := models.NewUser(userID)

	err := b.Database.SubscribeUser(user)
	if err != nil {
		return err
	}

	return nil
}
