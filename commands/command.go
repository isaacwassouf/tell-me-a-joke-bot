package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command interface {
	Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update)
}

func CreateCommand(command string, description string) *tgbotapi.BotCommand {
	return &tgbotapi.BotCommand{Command: command, Description: description}
}

func RegisterCommands(bot *tgbotapi.BotAPI, commands ...tgbotapi.BotCommand) {
	cfg := tgbotapi.NewSetMyCommands(commands...)
	bot.Send(cfg)
}

func CreateAndRegisterCommands(bot *tgbotapi.BotAPI) {
	RegisterCommands(bot,
		*CreateCommand("description", "Get the description of the bot"),
		*CreateCommand("joke", "Tell me a joke"),
		*CreateCommand("subscribe", "Subscribe to jokes"),
		*CreateCommand("unsubscribe", "Unsubscribe to jokes"),
	)
}
