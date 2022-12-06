package boltdb

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/cr00z/goTelegram/pkg/repository"
)

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Put(chatID int64, token string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

func (r *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(chatID))
		token = string(data)
		if token == "" {
			return errors.New("token not found")
		}
		return nil
	})

	return token, err
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
