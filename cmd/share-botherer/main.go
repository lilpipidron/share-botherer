package main

import (
	"github.com/joho/godotenv"
	"github.com/lilpipidron/share-botherer/internal/bot"
	"github.com/lilpipidron/share-botherer/internal/config"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := config.Load()
	db := postgresql.InitDB(cfg)
	tgBot := bot.Connect(cfg)

	tgBot.Handle("/start", bot.Start(tgBot, db))
	tgBot.Handle("/pair", bot.Pair(tgBot, db))
	tgBot.Handle("/send", bot.Send(tgBot, db))
	tgBot.Handle("/mark", bot.Mark(tgBot, db))
	tgBot.Handle("/help", bot.Help())
	go bot.Sender(tgBot, db)

	tgBot.Start()
}
