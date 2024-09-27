package main

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/lilpipidron/share-botherer/internal/bot"
	"github.com/lilpipidron/share-botherer/internal/config"
	"github.com/lilpipidron/share-botherer/internal/storage"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
	"github.com/lilpipidron/share-botherer/internal/storage/sqlite"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config := config.Load()
	tgBot := bot.Connect(config)

	storage := getStorage(config)

	tgBot.Handle("/start", bot.Start(tgBot, storage))
	tgBot.Handle("/pair", bot.Pair(tgBot, storage))
	tgBot.Handle("/send", bot.Send(tgBot, storage))
	tgBot.Handle("/mark", bot.Mark(tgBot, storage))
	tgBot.Handle("/help", bot.Help())
	go bot.Sender(tgBot, storage)

	log.Info("starting bot")
	tgBot.Start()
}

func getStorage(config *config.Config) storage.IStorage {
	if config.PostgresConfig.Empty {
		storage, err := sqlite.NewStorage()
		if err != nil {
			log.Fatal("failed to initialize sqlite storage", "err", err)
		}

		log.Info("initialized sqlite storage")
		return storage
	}

	storage, err := postgresql.NewStorage(config)
	if err != nil {
		log.Fatal("failed to initialize postgres storage", "err", err)
	}

	log.Info("initialized postgres storage")
	return storage
}
