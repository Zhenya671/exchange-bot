package telegram

import (
	"fmt"
	bnb "github.com/Zhenya671/go-bnb-sdk"
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
)

const (
	usdID  = 431
	euroID = 451
	gbpID  = 429
)

const (
	commandStart = "start"
	commandUSD   = "usd"
	commandEUR   = "eur"
	commandGBP   = "gbp"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	chatID := message.Chat.ID
	intRand := rand.Int63n(2)
	msg := tgbotapi.NewMessage(chatID, message.Text)

	if intRand == 1 {
		msg = tgbotapi.NewMessage(chatID, "Today is winner day. Try your luck in smth bud")
	}

	if err := b.dataUsers.Save(chatID, message.From.FirstName, repository.UserData); err != nil {
		return err
	}

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleCommandStart(message)
	case commandUSD:
		return b.handleCommandUSD(message)
	case commandEUR:
		return b.handleCommandEURO(message)
	case commandGBP:
		return b.handleCommandGBP(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Choose one of three command and get current rate..")

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommandUSD(message *tgbotapi.Message) error {
	currency, er := bnb.GetCurrentCurrency(usdID)
	if er != nil {
		return er
	}
	response := prepareResponse(currency)
	msg := prepareMessage(message.Chat.ID, response)

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommandEURO(message *tgbotapi.Message) error {
	currency, er := bnb.GetCurrentCurrency(euroID)
	if er != nil {
		return er
	}
	response := prepareResponse(currency)
	msg := prepareMessage(message.Chat.ID, response)

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommandGBP(message *tgbotapi.Message) error {
	currency, er := bnb.GetCurrentCurrency(gbpID)
	if er != nil {
		return er
	}
	response := prepareResponse(currency)
	msg := prepareMessage(message.Chat.ID, response)

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "i don't know what is command")

	_, err := b.bot.Send(msg)
	return err
}

func prepareResponse(currency map[string]interface{}) []interface{} {
	var response []interface{}
	for key, value := range currency {
		if key == "Cur_Abbreviation" {
			response = append(response, value)
		}
		if key == "Cur_Name" {
			response = append(response, value)
		}
		if key == "Cur_OfficialRate" {
			response = append(response, fmt.Sprintf("%v", value))
		}
	}

	return response
}

func prepareMessage(chatID int64, response []interface{}) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, fmt.Sprintf("%s - %s %v", response[0], response[1], response[2]))
}
