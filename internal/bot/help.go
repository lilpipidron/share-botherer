package bot

import (
	"gopkg.in/telebot.v3"
)

func Help() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := "/pair username - creates a pair with your friend. Both users " +
			"must enter this command for a successful connection\n" +
			"/send username text - sends a message to your friend. /pair is required beforehand\n" +
			"/mark - use this command and forward the message at the same time. " +
			"The message will be marked as read and will not be sent again"

		return c.Send(message)
	}
}
