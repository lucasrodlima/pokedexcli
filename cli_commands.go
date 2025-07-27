package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands = map[string]cliCommand{
	"help": {},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Show Area Locations",
		callback:    commandMap,
	},
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap() error {
	areas := locationAreas{}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area")

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error in poke api request: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&areas); err != nil {
		return fmt.Errorf("Error decoding json: %w", err)
	}

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}
	return nil
}
