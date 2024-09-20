package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/isaacwassouf/get-the-tee-and-mate/commands"
	"github.com/isaacwassouf/get-the-tee-and-mate/jokes"
)

type CommandHandler struct {
	Handlers       map[string]commands.Command
	JokeObservable *jokes.JokeObservable
}

func NewCommandHandler(JokeObservable *jokes.JokeObservable) *CommandHandler {
	handlers := make(map[string]commands.Command)
	handlers["description"] = &commands.DescriptionCommand{}
	handlers["joke"] = &commands.JokeCommand{}
	handlers["subscribe"] = &commands.SubscribeCommand{JokeObservable: JokeObservable}
	handlers["unsubscribe"] = &commands.UnsubscribeCommand{JokeObservable: JokeObservable}

	return &CommandHandler{
		Handlers:       handlers,
		JokeObservable: JokeObservable,
	}
}

func (h *CommandHandler) HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	rawCommand := update.Message.Command()
	if rawCommand == "" {
		return
	}

	handler, ok := h.Handlers[rawCommand]
	if !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You're trying to be smart, huh? 🤔")
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
		return
	}

	handler.Handle(bot, update)
}
