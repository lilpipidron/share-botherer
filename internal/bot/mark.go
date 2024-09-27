package bot

import (
	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/storage"
	"gopkg.in/telebot.v3"
)

func Mark(bot *telebot.Bot, storage storage.IStorage) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := c.Message()
		if !message.IsReply() {
			return c.Send("You must reply to the message you want to delete")
		}

		text := message.ReplyTo.Text

		if err := storage.DeleteMessage(text, message.Sender.ID); err != nil {
			log.Error(err)
			return c.Send("Failed to delete message")
		}

		return c.Send("Marked")
	}
}
