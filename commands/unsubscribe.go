package commands

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/isaacwassouf/get-the-tee-and-mate/jokes"
	"github.com/isaacwassouf/get-the-tee-and-mate/models"
)

type UnsubscribeCommand struct {
	JokeObservable *jokes.JokeObservable
}

func (c *UnsubscribeCommand) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	err := c.JokeObservable.Unsubscribe(&models.User{
		User:   *update.Message.From,
		ChatID: update.Message.Chat.ID,
	})
	if err != nil {
		if errors.Is(err, jokes.NotSubscribedError{}) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are not subscribed to the jokes! 🤔")
			bot.Send(msg)
			return
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You have been unsubscribed from the jokes! 😢")
	bot.Send(msg)
}
