package botmsgpipeline

import (
	"fmt"
	"os"
	"strings"
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

	channelIds []string

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

	pcChan    chan *discordgo.MessageCreate
	wgPc      *sync.WaitGroup
	pcWorkers int
}

func New(botClient IBotClient, stockAPI IStockAPI, db IDatabase, inbound chan *discordgo.MessageCreate, log *logging.Log) (*BotPipeline, error) {

	return &BotPipeline{
		botClient: botClient,
		stockAPI:  stockAPI,
		db:        db,
		logger:    log,

		channelIds: strings.Split(os.Getenv("channel_ids"), ","),

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

		pcChan:    make(chan *discordgo.MessageCreate),
		wgPc:      &sync.WaitGroup{},
		pcWorkers: 10,
	}, nil
}

func isValidTradingTime() bool {
	loc, err := time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
		return false
	}

	now := time.Now().In(loc)

	dayString := fmt.Sprintf("%d", now.Day())
	monthString := fmt.Sprintf("%d", now.Month())

	if now.Day() < 10 {
		dayString = fmt.Sprintf("0%d", now.Day())
	}

	if now.Month() < 10 {
		monthString = fmt.Sprintf("0%d", now.Month())
	}

	startString := fmt.Sprintf("%d-%s-%sT09:30:00.00-05:00", now.Year(), monthString, dayString)
	endString := fmt.Sprintf("%d-%s-%sT16:00:00.00-05:00", now.Year(), monthString, dayString)

	start, err := time.Parse(time.RFC3339, startString)
	if err != nil {
		fmt.Println(err)
		return false
	}
	end, err := time.Parse(time.RFC3339, endString)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// If saturday or sunday
	if now.Weekday() == 6 || now.Weekday() == 0 {
		return false
	}

	fmt.Println(now)
	fmt.Println(end)
	fmt.Println(now.After(start), now.Before(end.In(loc)))

	if now.After(start) && now.Before(end) {
		return true
	}

	return false
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
