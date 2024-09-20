package models

type Joke struct {
	JokeType      string `json:"type"`
	JokeSetup     string `json:"setup"`
	JokePunchline string `json:"punchline"`
	JokeID        int    `json:"id"`
}
