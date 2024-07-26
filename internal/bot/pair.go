package bot

import (
	"errors"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

func Pair(bot *telebot.Bot, storage *postgresql.StorageGorm) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := c.Message().Text
		words := strings.Split(message, " ")
		if len(words) != 2 {
			return c.Send("Command struct: /pair username")
		}

		username := words[1]
		pair := &models.UserConnection{}
		currentUserID := c.Message().Sender.ID
		pairUser := &models.User{}

		if err := storage.DB.First(pairUser, "username = ?", username).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Info("Not found", "username", username)
			} else {
				log.Error(err)
			}
			return c.Send("User not found. Make sure the user is registered in the bot.")
		}

		if err := storage.DB.First(pair, "user_id1 = ? and user_id2 = ?", pairUser.TelegramID, currentUserID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Error(err)
				return c.Send("Pairing failed")
			}
			pair.UserID1 = currentUserID
			pair.UserID2 = pairUser.TelegramID

			if err := storage.DB.Save(pair).Error; err != nil {
				log.Error(err)
				return c.Send("Pairing failed")
			}

			return c.Send("Pairing request created")
		}

		pair.Paired = true

		if err := storage.DB.Save(pair).Error; err != nil {
			log.Error(err)
			return c.Send("Pairing failed")
		}

		return c.Send("Paired")
	}
}

func isNotFound(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == pgerrcode.NoDataFound
	}
	return false
}
