package bot

import (
	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/config"
	"gopkg.in/telebot.v3"
)

func Connect(cfg *config.Config) *telebot.Bot {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: cfg.Token,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Info("connected to bot")

	return bot
}
