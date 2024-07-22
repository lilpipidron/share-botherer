package bot

import (
	"github.com/charmbracelet/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lilpipidron/share-botherer/internal/config"
)

func Start(cfg *config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.Toket)
	if err != nil {
		panic(err)
	}

	log.Debug("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for range cfg.WorkersForUpdates {
		go worker(updates, bot)
	}
}

func worker(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message != nil {
			log.Info("Message received", "username", update.Message.From.UserName,
				"chatID", update.Message.Chat.ID, "message text", update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "test message")

			if _, err := bot.Send(msg); err != nil {
				log.Error(err)
			}
		}
	}
}
