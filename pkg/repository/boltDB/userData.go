package boltDB

import (
	"errors"
	"github.com/Zhenya671/telegram-bot-exchangeRates/pkg/repository"
	"github.com/boltdb/bolt"
	"strconv"
)

type UserDataRepository struct {
	db *bolt.DB
}

func NewUserDataRepository(db *bolt.DB) *UserDataRepository {
	return &UserDataRepository{db: db}
}

func (r *UserDataRepository) Save(chatID int64, firstName string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToByte(chatID), []byte(firstName))
	})
}

func (r *UserDataRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var name string
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToByte(chatID))
		name = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}

	if name == "" {
		return "", errors.New("name not found")
	}
	return name, nil
}

func intToByte(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
