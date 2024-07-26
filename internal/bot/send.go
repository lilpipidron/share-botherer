package bot

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
	"gopkg.in/telebot.v3"
)

func Send(bot *telebot.Bot, storage *postgresql.StorageGorm) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := c.Message().Text
		words := strings.Split(message, " ")
		if len(words) < 3 {
			return c.Send("Command struct: /send username message...")
		}

		username := words[1]
		pair := &models.UserConnection{}
		currentUserID := c.Message().Sender.ID
		receiver := &models.User{}

		if err := storage.DB.First(receiver, "username = ?", username).Error; err != nil {
			log.Error(err)
			return c.Send("User not found. Make sure the user is registered in the bot.")
		}

		if err := storage.DB.First(pair, "user_id1 = ? and user_id2 = ? or user_id1 = ? and user_id2 = ?", currentUserID, receiver.TelegramID, receiver.TelegramID, currentUserID).Error; err != nil {
			log.Error(err)
			return c.Send("First you have to create a pair with the user using command /pair")
		}

		m := &models.Message{
			FromUserID: currentUserID,
			ToUserID:   receiver.TelegramID,
			Text:       strings.Join(words[2:], " "),
		}

		m.DeleteKey = generateMessageHash(m.FromUserID, m.ToUserID, m.Text)

		if err := storage.DB.Save(m).Error; err != nil {
			log.Error(err)
			return c.Send("Failed to send")
		}

		return c.Send("Sended")
	}
}

func generateMessageHash(fromUserID, toUserID int64, text string) string {
	data := []byte(strconv.Itoa(int(fromUserID)) + strconv.Itoa(int(toUserID)) + text)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
