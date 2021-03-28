package botmsgpipeline

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	Prefix    = "$"
	BuyShares = "BUY"
	BuyMoney  = "BUYMONEY"
	Sell      = "SELL"
	Join      = "JOIN"
	Stats     = "STATS"
	Help      = "HELP"
	Pc        = "PC"

	stonksChannelID    = "804809948583559220"
	stonksDevChannelID = "804911890484428830"
)

func (b *BotPipeline) Start() {

	for i := 0; i < b.validateWorkers; i++ {
		go b.validate()
		b.wgValidate.Add(1)
	}

	for i := 0; i < b.parseWorkers; i++ {
		go b.parse()
		b.wgParse.Add(1)
	}

	for i := 0; i < b.buySharesWorkers; i++ {
		go b.buyShares()
		b.wgBuyShares.Add(1)
	}

	for i := 0; i < b.buyMoneyWorkers; i++ {
		go b.buyMoney()
		b.wgBuyMoney.Add(1)
	}

	for i := 0; i < b.sellSharesWorkers; i++ {
		go b.sellShares()
		b.wgSellShares.Add(1)
	}

	for i := 0; i < b.joinWorkers; i++ {
		go b.join()
		b.wgJoin.Add(1)
	}

	for i := 0; i < b.statsWorkers; i++ {
		go b.stats()
		b.wgStats.Add(1)
	}

	for i := 0; i < b.helpWorkers; i++ {
		go b.help()
		b.wgHelp.Add(1)
	}

	for i := 0; i < b.pcWorkers; i++ {
		go b.priceCheck()
		b.wgPc.Add(1)
	}

}

func (b *BotPipeline) validate() {
	for msg := range b.validateChan {
		if !stringInSlice(msg.ChannelID, b.channelIds) {
			fmt.Println(msg.ChannelID, b.channelIds)
			continue
		}

		if len(msg.Content) < 1 {
			continue
		}

		if string(msg.Content[0]) == Prefix {
			parsed := msg.Content[1:len(msg.Content)]
			tokenized := strings.Split(parsed, " ")
			if len(tokenized) < 1 {
				continue
			}
			b.parseChan <- msg
		}

	}
}

func (b *BotPipeline) parse() {
	for msg := range b.parseChan {
		parsed := msg.Content[1:len(msg.Content)]
		tokenized := strings.Split(parsed, " ")
		command := tokenized[0]

		switch strings.ToUpper(command) {
		case BuyShares:
			b.buySharesChan <- msg

		case BuyMoney:
			b.buyMoneyChan <- msg

		case Sell:
			b.sellSharesChan <- msg

		case Join:
			b.joinChan <- msg

		case Stats:
			b.statsChan <- msg

		case Help:
			b.helpChan <- msg

		case Pc:
			b.pcChan <- msg
		}
	}
}

func extractTickerQuantity(tokenized []string) (string, float64, error) {
	if len(tokenized) < 3 {
		return "", 0, errors.New("Not enough args")
	}
	ticker := strings.ToUpper(tokenized[1])
	quantity := tokenized[2]

	quantityFloat, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return "", 0, err
	}

	return ticker, quantityFloat, nil
}
