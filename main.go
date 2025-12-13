package main

import (
	"time"

	"github.com/CodeZeroSugar/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	NextURL       *string
	PreviousURL   *string
}

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeapiClient: client,
		NextURL:       nil,
		PreviousURL:   nil,
	}

	startRepl(cfg)
}
