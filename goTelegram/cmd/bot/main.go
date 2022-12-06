package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/cr00z/goTelegram/pkg/repository"
	"github.com/cr00z/goTelegram/pkg/repository/boltdb"
	"github.com/cr00z/goTelegram/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5417402887:AAEaytCQ5wHHNPCofKx_RkHVtR7afE8PfQ0")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	// pocketClient, err := pocket.NewClient("103982-153103b17c3e7f25d958959")
	pocketClient, err := pocket.NewClient("103995-c1849b95484073bb93b76e7")
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
