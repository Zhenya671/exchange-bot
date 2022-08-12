package main

import (
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/repository/boltDB"
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5558699531:AAHIaGPR4utxz4yz5WYtqUO5NWEQnHFuBVQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	db, err := bolt.Open("db/bot.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	userDataRepository := boltDB.NewUserDataRepository(db)

	telegramBot := telegram.NewBot(bot, userDataRepository)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
