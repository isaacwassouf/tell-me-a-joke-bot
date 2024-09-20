package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/isaacwassouf/get-the-tee-and-mate/commands"
	"github.com/isaacwassouf/get-the-tee-and-mate/jokes"
	"github.com/isaacwassouf/get-the-tee-and-mate/utils"
)

func main() {
	// load the environment variables from the .env file
	err := utils.LoadEnvFromFile()
	if err != nil {
		log.Fatal(err)
	}

	// get the token from the environment
	apiToken, err := utils.GetEnvVar("TELEGRAM_API_TOKEN", "")
	if err != nil {
		log.Fatal(err)
	}

	// create a new bot
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// register the commands
	commands.CreateAndRegisterCommands(bot)

	// create a new observable to fetch jokes
	jocksObservable := jokes.NewJokeObervable(bot)
	// create the command handler
	commandHandler := utils.NewCommandHandler(jocksObservable)

	// fetch the jokes from the observable
	go jocksObservable.FetchJokes()

	updater := tgbotapi.NewUpdate(0)
	updater.Timeout = 60
	updates := bot.GetUpdatesChan(updater)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				commandHandler.HandleCommand(bot, update)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What're you trying to play at lil' buddy? 🤨 I only understand commands")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}

// func getQuote() (string, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
//
// 	respch := make(chan *http.Response)
//
// 	go func() {
// 		res, err := http.Get("https://zenquotes.io/api/random")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		respch <- res
// 	}()
//
// 	select {
// 	case res := <-respch:
// 		defer res.Body.Close()
// 		data, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return string(data), nil
// 	case <-ctx.Done():
// 		return "", ctx.Err()
//
// 	}
// }

// func quoteLooper(qchan chan string) {
// 	// get a quote every 10 seconds
// 	ticker := time.NewTicker(10 * time.Second)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			quote, err := getQuote()
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			qchan <- quote
// 		}
// 	}
// }
