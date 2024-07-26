package bot

import (
	"errors"

	"github.com/charmbracelet/log"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
	"gopkg.in/telebot.v3"
)

func Start(bot *telebot.Bot, storage *postgresql.StorageGorm) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		user := models.User{
			ChatID:     c.Chat().ID,
			TelegramID: c.Message().Sender.ID,
		}

		if err := storage.DB.Save(&user).Error; err != nil {
			log.Error(err)
			if isUniqueViolation(err) {
				return c.Send("User already exists")
			}
			return err
		}

		return c.Send("User saved")
	}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == pgerrcode.UniqueViolation
	}
	return false
}
