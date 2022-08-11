package main

import (
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5558699531:AAHIaGPR4utxz4yz5WYtqUO5NWEQnHFuBVQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
