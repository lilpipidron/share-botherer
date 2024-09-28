package bot

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage"
	"gopkg.in/telebot.v3"
)

func Send(bot *telebot.Bot, storage storage.IStorage) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := c.Message().Text
		words := strings.Split(message, " ")
		if len(words) < 3 {
			return c.Send("Command: /send username message...")
		}

		username := words[1]
		username = strings.TrimPrefix(username, "@")
		currentUserID := c.Message().Sender.ID
		receiver, err := storage.FindUserByUsername(username)
		if err != nil {
			log.Error(err)
			return c.Send(
				"User not found. Make sure the user is registered in the bot.")
		}

		_, err = storage.FindUserConnection(currentUserID, receiver.TelegramID)
		if err != nil {
			log.Error(err)
			return c.Send(
				"First you have to create a pair with the user using command /pair")
		}

		newMessage := &models.Message{
			FromUserID: currentUserID,
			ToUserID:   receiver.TelegramID,
			Text:       strings.Join(words[2:], " "),
		}

		if err := storage.SaveMessage(newMessage); err != nil {
			log.Error(err)
			return c.Send("Failed to send")
		}

		return c.Send("Sended")
	}
}
