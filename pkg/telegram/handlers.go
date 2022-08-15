package telegram

import (
	"fmt"
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	commandStart    = "start"
	commandRetrieve = "retrieve"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	chatID := message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, message.Text)

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
	case commandRetrieve:
		return b.handleCommandRetrieve(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "command start")

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommandRetrieve(message *tgbotapi.Message) error {
	collectMsg, err := b.dataUsers.Get(message.Chat.ID, repository.UserData)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("\n%s", collectMsg))

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "i don't know what is command")

	_, err := b.bot.Send(msg)
	return err
}
