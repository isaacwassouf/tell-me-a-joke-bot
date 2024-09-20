package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DescriptionCommand struct{}

func (c *DescriptionCommand) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I'm a bot that tells you a joke every 1 hour, or whenever you ask to :)")
	bot.Send(msg)
}
