package botmsgpipeline

import (
	"fmt"
	"sync"
	"time"

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

	validateChan    chan *discordgo.MessageCreate
	wgValidate      *sync.WaitGroup
	validateWorkers int

	parseChan    chan *discordgo.MessageCreate
	wgParse      *sync.WaitGroup
	parseWorkers int

	buySharesChan    chan *discordgo.MessageCreate
	wgBuyShares      *sync.WaitGroup
	buySharesWorkers int

	buyMoneyChan    chan *discordgo.MessageCreate
	wgBuyMoney      *sync.WaitGroup
	buyMoneyWorkers int

	sellSharesChan    chan *discordgo.MessageCreate
	wgSellShares      *sync.WaitGroup
	sellSharesWorkers int

	statsChan    chan *discordgo.MessageCreate
	wgStats      *sync.WaitGroup
	statsWorkers int

	joinChan    chan *discordgo.MessageCreate
	wgJoin      *sync.WaitGroup
	joinWorkers int

	helpChan    chan *discordgo.MessageCreate
	wgHelp      *sync.WaitGroup
	helpWorkers int
}

func New(botClient IBotClient, stockAPI IStockAPI, db IDatabase, inbound chan *discordgo.MessageCreate, log *logging.Log) (*BotPipeline, error) {

	return &BotPipeline{
		botClient: botClient,
		stockAPI:  stockAPI,
		db:        db,
		logger:    log,

		validateChan:    inbound,
		wgValidate:      &sync.WaitGroup{},
		validateWorkers: 10,

		parseChan:    make(chan *discordgo.MessageCreate),
		wgParse:      &sync.WaitGroup{},
		parseWorkers: 10,

		buySharesChan:    make(chan *discordgo.MessageCreate),
		wgBuyShares:      &sync.WaitGroup{},
		buySharesWorkers: 10,

		buyMoneyChan:    make(chan *discordgo.MessageCreate),
		wgBuyMoney:      &sync.WaitGroup{},
		buyMoneyWorkers: 10,

		sellSharesChan:    make(chan *discordgo.MessageCreate),
		wgSellShares:      &sync.WaitGroup{},
		sellSharesWorkers: 10,

		statsChan:    make(chan *discordgo.MessageCreate),
		wgStats:      &sync.WaitGroup{},
		statsWorkers: 10,

		joinChan:    make(chan *discordgo.MessageCreate),
		wgJoin:      &sync.WaitGroup{},
		joinWorkers: 10,

		helpChan:    make(chan *discordgo.MessageCreate),
		wgHelp:      &sync.WaitGroup{},
		helpWorkers: 10,
	}, nil
}

func isValidTradingTime() bool {
	loc, err := time.LoadLocation("NYC")
	if err != nil {
		return false
	}

	now := time.Now().In(loc)

	startString := fmt.Sprintf("%d-%d-%dT09:30:00.00Z", now.Year(), now.Month(), now.Day())
	endString := fmt.Sprintf("%d-%d-%dT04:00:00.00Z", now.Year(), now.Month(), now.Day())

	start, err := time.Parse(time.RFC3339, startString)
	if err != nil {
		return false
	}
	end, err := time.Parse(time.RFC3339, endString)
	if err != nil {
		return false
	}

	// If saturday or sunday
	if now.Weekday() == 6 || now.Weekday() == 0 {
		return false
	}

	if now.After(start) && now.Before(end) {
		return true
	}

	return false
}
