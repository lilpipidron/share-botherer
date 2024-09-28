package bot

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

func Pair(bot *telebot.Bot, storage storage.IStorage) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := c.Message().Text
		words := strings.Split(message, " ")
		if len(words) != 2 {
			return c.Send("Command struct: /pair username")
		}

		username := words[1]
		username = strings.TrimPrefix(username, "@")
		pair := &models.UserConnection{}
		currentUserID := c.Message().Sender.ID
		pairUser, err := storage.FindUserByUsername(username)
		if err != nil {
			log.Error(err)
			return c.Send("User not found. Make sure the user is registered in the bot.")
		}

		pair, err = storage.FindUserConnection(pairUser.TelegramID, currentUserID)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				log.Error(err)
				return c.Send("Pairing failed")
			}
			pair.UserID1 = currentUserID
			pair.UserID2 = pairUser.TelegramID

			if err := storage.SaveUserConnection(pair); err != nil {
				log.Error(err)
				return c.Send("Pairing failed")
			}

			return c.Send("Pairing request created")
		}

		pair.Paired = true

		if err := storage.SaveUserConnection(pair); err != nil {
			log.Error(err)
			return c.Send("Pairing failed")
		}

		return c.Send("Paired")
	}
}
