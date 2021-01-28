package bot

import (
	"github.com/J-Rivard/trading-bot/internal/logging"
	"github.com/J-Rivard/trading-bot/internal/models"
	"github.com/bwmarrin/discordgo"
)

type IStockAPI interface {
	GetStockData(ticker string) (*models.Stock, error)
}

type IDatabase interface {
	SubscribeUser(user *models.User) error
	GetUser(id string) (*models.User, error)
}

type Bot struct {
	Client   *discordgo.Session
	Log      *logging.Log
	Database IDatabase
	stockAPI IStockAPI
}

type Parameters struct {
	Token string
}

const (
	stonksChannelID = "804114998275604511"
)

func New(params *Parameters, stockAPI IStockAPI, database IDatabase, log *logging.Log) (*Bot, error) {
	dg, err := discordgo.New("Bot " + params.Token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Client:   dg,
		Log:      log,
		Database: database,
		stockAPI: stockAPI,
	}, nil
}

func (b *Bot) Start() error {
	b.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err := b.Client.Open()
	if err != nil {
		return err
	}

	err = b.SetupHandlers(b.messageCreate)

	return err
}

func (b *Bot) Stop() error {
	err := b.Client.Close()
	return err
}

func (b *Bot) SetupHandlers(handlers ...interface{}) error {
	for _, handler := range handlers {
		b.Client.AddHandler(handler)
	}

	return nil
}
