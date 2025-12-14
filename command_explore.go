package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location")
	}
	locationName := args[0]
	fmt.Printf("Exploring %v...\n", locationName)

	details, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, enc := range details.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
