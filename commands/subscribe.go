package commands

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/isaacwassouf/get-the-tee-and-mate/jokes"
	"github.com/isaacwassouf/get-the-tee-and-mate/models"
)

type SubscribeCommand struct {
	JokeObservable *jokes.JokeObservable
}

func (c *SubscribeCommand) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	err := c.JokeObservable.Subscribe(&models.User{
		User:   *update.Message.From,
		ChatID: update.Message.Chat.ID,
	})
	if err != nil {
		if errors.Is(err, jokes.SubscriberExistsError{}) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are already subscribed to the jokes! 🤔")
			bot.Send(msg)
			return
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You have been subscribed to the jokes! 🎉")
	bot.Send(msg)
}
