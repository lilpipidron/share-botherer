package bot

import (
	"gopkg.in/telebot.v3"
)

func Help() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		message := "/pair username - creates a pair with your friend. Both of you " +
			"should enter this command for a connection\n" +
			"/send username text - sends a message to your friend. /pair is required beforehand\n" +
			"/mark - use this command while replying to the message you want marked. " +
			"The message will be marked as read and will not be sent again"

		return c.Send(message)
	}
}
