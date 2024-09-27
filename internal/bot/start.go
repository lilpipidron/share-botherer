package bot

import (
	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage"
	"gopkg.in/telebot.v3"
)

func Start(bot *telebot.Bot, storage storage.IStorage) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		user := models.User{
			ChatID:     c.Chat().ID,
			TelegramID: c.Message().Sender.ID,
			Username:   c.Chat().Username,
		}

		if err := storage.SaveUser(&user); err != nil {
			log.Error(err)
			return err
		}

		return c.Send("To find out the list of available commands, type /help")
	}
}
