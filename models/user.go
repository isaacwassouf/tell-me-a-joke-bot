package models

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Observer interface {
	Notify(joke Joke)
}

type User struct {
	tgbotapi.User
	ChatID int64
}

func (u *User) Notify(bot *tgbotapi.BotAPI, joke Joke) {
	msgText := fmt.Sprintf("<b>%s</b> \n <i>%s</i>", joke.JokeSetup, joke.JokePunchline)
	msg := tgbotapi.NewMessage(u.ChatID, msgText)
	msg.ParseMode = "HTML"

	bot.Send(msg)
}
