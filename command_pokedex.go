package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.ImprisonedCreatures) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}
	for _, name := range cfg.ImprisonedCreatures {
		fmt.Printf("- %s\n", name.Name)
	}

	return nil
}
