package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for name, command := range getCommands() {
		fmt.Printf("%v: %v\n", name, command.description)
	}
	return nil
}
