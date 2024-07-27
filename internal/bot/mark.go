package bot

import (
	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
	"gopkg.in/telebot.v3"
)

func Mark(bot *telebot.Bot, storage *postgresql.StorageGorm) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := c.Message()
		if !message.IsReply() {
			return c.Send("You must reply to the message you want to delete")
		}

		text := message.ReplyTo.Text

		if err := storage.DB.Delete(models.Message{}, "text = ? and to_user_id = ?", text, message.Sender.ID).Error; err != nil {
			log.Error(err)
			return c.Send("Failed to delete message")
		}

		return c.Send("Marked")
	}
}
