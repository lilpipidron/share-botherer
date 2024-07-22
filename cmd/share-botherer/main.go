package main

import (
	"github.com/joho/godotenv"
	"github.com/lilpipidron/share-botherer/internal/bot"
	"github.com/lilpipidron/share-botherer/internal/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := config.Load()
	bot.Start(cfg)
}
