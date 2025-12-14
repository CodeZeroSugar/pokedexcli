package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a Pokemon name")
	}
	pokeName := args[0]

	_, exists := cfg.ImprisonedCreatures[pokeName]
	if exists {
		fmt.Printf("%s has already been caught...\n", pokeName)
		return nil
	}

	pokeInfo, err := cfg.pokeapiClient.GetPokeInfo(pokeName)
	if err != nil {
		fmt.Println("please enter a real Pokemon name...")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)

	maxExp := 255.0
	exp := float64(pokeInfo.BaseExperience)
	if exp > maxExp {
		exp = maxExp
	}
	chance := 1.0 - (exp / maxExp)
	roll := rand.Float64()

	if roll < chance {
		fmt.Printf("%s was caught!\n", pokeName)
		fmt.Println("You may now inspect it with the 'inspect' command.")
	} else {
		fmt.Printf("%s got away!\n", pokeName)
		return nil
	}

	cfg.ImprisonedCreatures[pokeName] = pokeInfo

	return nil
}
