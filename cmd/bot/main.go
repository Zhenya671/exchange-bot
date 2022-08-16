package main

import (
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/config"
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/repository"
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/repository/boltDB"
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	cnf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cnf)

	bot, err := tgbotapi.NewBotAPI(cnf.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	db, err := initDB(cnf)
	if err != nil {
		log.Fatal(err)
	}

	userDataRepository := boltDB.NewUserDataRepository(db)

	telegramBot := telegram.NewBot(bot, userDataRepository)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB(cnf *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cnf.DBPath, 0600, nil)
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
