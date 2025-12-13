package main

import (
	"bufio"
	"encoding/json"
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
	callback    func(*config) error
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
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for name, command := range getCommands() {
		fmt.Printf("%v: %v\n", name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	nextURL := cfg.NextURL
	body, err := cfg.pokeapiClient.Fetch(*nextURL)
	if err != nil {
		return fmt.Errorf("error! 'Fetch' method failed: %v", err)
	}

	locationAreaResp := LocationAreasResponse{}
	err = json.Unmarshal(body, &locationAreaResp)
	if err != nil {
		return fmt.Errorf("error! Failed to Unmarshal json: %v", err)
	}

	for _, locationResults := range locationAreaResp.Results {
		fmt.Println(locationResults.Name)
	}
	if locationAreaResp.Next != nil {
		cfg.NextURL = locationAreaResp.Next
	} else {
		cfg.NextURL = nil
	}

	if locationAreaResp.Previous != nil {
		cfg.PreviousURL = locationAreaResp.Previous
	} else {
		cfg.PreviousURL = nil
	}

	return nil
}

func commandMapb(cfg *config) error {
	prevURL := cfg.PreviousURL
	if cfg.PreviousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	body, err := cfg.pokeapiClient.Fetch(*prevURL)
	if err != nil {
		return fmt.Errorf("error! 'Fetch' method failed: %v", err)
	}

	locationAreaResp := LocationAreasResponse{}
	err = json.Unmarshal(body, &locationAreaResp)
	if err != nil {
		return fmt.Errorf("error! Failed to Unmarshal json: %v", err)
	}

	for _, locationResults := range locationAreaResp.Results {
		fmt.Println(locationResults.Name)
	}
	if locationAreaResp.Next != nil {
		cfg.NextURL = locationAreaResp.Next
	} else {
		cfg.NextURL = nil
	}

	if locationAreaResp.Previous != nil {
		cfg.PreviousURL = locationAreaResp.Previous
	} else {
		cfg.PreviousURL = nil
	}

	return nil
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
		commandList := getCommands()

		if commandStruct, exists := commandList[commandName]; exists {
			err := commandStruct.callback(cfg)
			if err != nil {
				fmt.Printf("Error occured during command callback: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
