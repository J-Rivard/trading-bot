package botmsgpipeline

import (
	"sync"

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
	UpdateUser(user *models.User) error
}

type IBotClient interface {
	SendMessage(id, msg string) (*discordgo.Message, error)
}

type BotPipeline struct {
	botClient IBotClient
	stockAPI  IStockAPI
	db        IDatabase
	logger    *logging.Log

	validateChan chan *discordgo.MessageCreate
	wgValidate   *sync.WaitGroup

	parseChan chan *discordgo.MessageCreate
	wgParse   *sync.WaitGroup

	buySharesChan chan *discordgo.MessageCreate
	wgBuyShares   *sync.WaitGroup

	buyMoneyChan chan *discordgo.MessageCreate
	wgBuyMoney   *sync.WaitGroup

	sellSharesChan chan *discordgo.MessageCreate
	wgSellShares   *sync.WaitGroup

	statsChan chan *discordgo.MessageCreate
	wgStats   *sync.WaitGroup

	joinChan chan *discordgo.MessageCreate
	wgJoin   *sync.WaitGroup

	helpChan chan *discordgo.MessageCreate
	wgHelp   *sync.WaitGroup
}

func New(botClient IBotClient, stockAPI IStockAPI, db IDatabase, inbound chan *discordgo.MessageCreate, log *logging.Log) (*BotPipeline, error) {

	return &BotPipeline{
		botClient: botClient,
		stockAPI:  stockAPI,
		db:        db,
		logger:    log,

		validateChan: inbound,
		wgValidate:   &sync.WaitGroup{},

		parseChan: make(chan *discordgo.MessageCreate),
		wgParse:   &sync.WaitGroup{},

		buySharesChan: make(chan *discordgo.MessageCreate),
		wgBuyShares:   &sync.WaitGroup{},

		buyMoneyChan: make(chan *discordgo.MessageCreate),
		wgBuyMoney:   &sync.WaitGroup{},

		sellSharesChan: make(chan *discordgo.MessageCreate),
		wgSellShares:   &sync.WaitGroup{},

		statsChan: make(chan *discordgo.MessageCreate),
		wgStats:   &sync.WaitGroup{},

		joinChan: make(chan *discordgo.MessageCreate),
		wgJoin:   &sync.WaitGroup{},

		helpChan: make(chan *discordgo.MessageCreate),
		wgHelp:   &sync.WaitGroup{},
	}, nil
}
