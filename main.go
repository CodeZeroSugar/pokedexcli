package main

import (
	"time"

	"github.com/CodeZeroSugar/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	NextURL             *string
	PreviousURL         *string
	ImprisonedCreatures map[string]pokeapi.PokeInfo
}

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	start := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	cfg := &config{
		pokeapiClient:       client,
		NextURL:             &start,
		PreviousURL:         nil,
		ImprisonedCreatures: make(map[string]pokeapi.PokeInfo),
	}

	startRepl(cfg)
}
