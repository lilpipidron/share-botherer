package bot

import (
	"time"

	"github.com/lilpipidron/share-botherer/internal/models"
	"github.com/lilpipidron/share-botherer/internal/storage/postgresql"
	"golang.org/x/exp/rand"
	"gopkg.in/telebot.v3"
)

func Sender(bot *telebot.Bot, storage *postgresql.StorageGorm) {
	for {
		randomDuration := time.Duration(rand.Intn(24*60*60)) * time.Second
		time.Sleep(randomDuration)

		var messages []models.Message

		sql := `
SELECT *
FROM (
    SELECT *, ROW_NUMBER() OVER(PARTITION BY from_user_id, to_user_id ORDER BY RANDOM() DESC) AS rn
    FROM messages
) t
WHERE rn = 1
ORDER BY RANDOM() DESC;
`

		storage.DB.Raw(sql).Scan(&messages)

		for _, message := range messages {
			bot.Send(&telebot.User{ID: message.ToUserID}, message.Text)
		}
	}
}
