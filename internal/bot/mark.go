package bot

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
	"gopkg.in/telebot.v3"
)

func Mark(bot *telebot.Bot, storage *postgresql.StorageGorm) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := c.Message().Text
		words := strings.Split(message, " ")
		if len(words) != 2 {
			return c.Send("Command struct: /mark delete key")
		}

		deleteKey := words[1]

		if err := storage.DB.Delete(models.Message{}, "delete_key = ?", deleteKey).Error; err != nil {
			log.Error(err)
			return c.Send("Failed to delete message")
		}

		return c.Send("Marked")
	}
}
