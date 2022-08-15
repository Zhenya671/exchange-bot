package telegram

import "errors"

var (
	errUnknownCommand = errors.New("unknown command")
)

//msg := tgbotapi.NewMessage(message.Chat.ID, "i don't know what is command")

func (b *Bot) handleError(chatID int64, err error) {
	switch err {

	}
}
