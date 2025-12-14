package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a Pokemon name")
	}
	pokeName := args[0]
	pokeStats, exists := cfg.ImprisonedCreatures[pokeName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokeStats.Name)
	fmt.Printf("Height: %d\n", pokeStats.Height)
	fmt.Printf("Weight: %d\n", pokeStats.Weight)
	fmt.Println("Stats")
	for _, stats := range pokeStats.Stats {
		fmt.Printf("-%s: %d\n", stats.Stat.Name, stats.Value)
	}
	fmt.Println("Types:")
	for _, types := range pokeStats.Types {
		fmt.Printf("- %s\n", types.Type.Name)
	}

	return nil
}
