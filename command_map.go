package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
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

func commandMapb(cfg *config, args ...string) error {
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
