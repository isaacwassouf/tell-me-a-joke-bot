package jokes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/isaacwassouf/get-the-tee-and-mate/models"
)

type JokeObservable struct {
	subscribers map[int64]*models.User
	bot         *tgbotapi.BotAPI
}

func NewJokeObervable(bot *tgbotapi.BotAPI) *JokeObservable {
	return &JokeObservable{
		bot:         bot,
		subscribers: make(map[int64]*models.User),
	}
}

func (j *JokeObservable) Subscribe(u *models.User) error {
	if _, ok := j.subscribers[u.ID]; ok {
		return SubscriberExistsError{}
	}

	j.subscribers[u.ID] = u
	return nil
}

func (j *JokeObservable) Unsubscribe(u *models.User) error {
	if _, ok := j.subscribers[u.ID]; !ok {
		return NotSubscribedError{}
	}

	delete(j.subscribers, u.ID)
	return nil
}

func (j *JokeObservable) SubscribersCount() int {
	return len(j.subscribers)
}

func (j *JokeObservable) Notify(joke models.Joke) {
	for _, subscriber := range j.subscribers {
		go subscriber.Notify(j.bot, joke)
	}
}

func (j *JokeObservable) FetchJokes() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			if j.SubscribersCount() > 0 {
				joke, err := GetJoke()
				if err != nil {
					log.Print("Error fetching the joke from the API", err.Error())
				} else {
					log.Printf("fetched the joke %+v from the API", joke)
					j.Notify(joke)
				}
			} else {
				log.Print("No subscribers to notify, not fetching a new joke")
			}
		}
	}
}

func GetJoke() (models.Joke, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	respch := make(chan *http.Response)

	go func() {
		res, err := http.Get("https://official-joke-api.appspot.com/jokes/random")
		if err != nil {
			log.Fatal(err)
		}
		respch <- res
	}()

	select {
	case res := <-respch:
		defer res.Body.Close()
		var joke models.Joke
		err := json.NewDecoder(res.Body).Decode(&joke)
		if err != nil {
			log.Fatal(err)
		}
		return joke, nil
	case <-ctx.Done():
		return models.Joke{}, ctx.Err()
	}
}
