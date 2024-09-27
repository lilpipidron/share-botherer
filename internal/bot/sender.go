package bot

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/lilpipidron/share-botherer/internal/storage"
	"golang.org/x/exp/rand"
	"gopkg.in/telebot.v3"
)

func Sender(bot *telebot.Bot, storage storage.IStorage) {
	for {
		randomDuration := time.Duration(rand.Intn(24*60*60)) * time.Second
		time.Sleep(randomDuration)

		messages, err := storage.GetRandomMessages()
		if err != nil {
			log.Error(err)
			continue
		}

		for _, message := range messages {
			bot.Send(&telebot.User{ID: message.ToUserID}, message.Text)
		}
	}
}
