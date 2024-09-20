package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/isaacwassouf/get-the-tee-and-mate/jokes"
)

type JokeCommand struct{}

func (c *JokeCommand) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	joke, err := jokes.GetJoke()
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I'm sorry, I couldn't get a joke for you 😔")
		bot.Send(msg)
		return
	}

	msgText := fmt.Sprintf("<b>%s</b> \n <i>%s</i>", joke.JokeSetup, joke.JokePunchline)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	msg.ParseMode = "HTML"

	bot.Send(msg)
}
