package main

import (
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/repository"
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

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	userDataRepository := boltDB.NewUserDataRepository(db)

	telegramBot := telegram.NewBot(bot, userDataRepository)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("db/bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.UserData))
		if err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}
