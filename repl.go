package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LocationAreasResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationArea
}

type LocationArea struct {
	Name string
	URL  string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area for Pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inpect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	cleanText := strings.TrimSpace(strings.ToLower(text))

	return strings.Fields(cleanText)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandArgs := words[1:]
		commandList := getCommands()

		if commandStruct, exists := commandList[commandName]; exists {
			err := commandStruct.callback(cfg, commandArgs...)
			if err != nil {
				fmt.Printf("Error occured during command callback: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
