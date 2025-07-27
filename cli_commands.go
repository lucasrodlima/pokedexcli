package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

var actualConfig *config

var cliCommands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "List Location Areas",
		callback:    commandMap,
		config:      actualConfig,
	},
	"mapb": {
		name:        "mapb",
		description: "List Previous Location Areas",
		callback:    commandMapB,
		config:      actualConfig,
	},
}

var helpMessage string

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	fmt.Println(helpMessage)
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap() error {
	areas := locationAreas{}
	var url string

	if actualConfig == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = actualConfig.next
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error in poke api request: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&areas); err != nil {
		return fmt.Errorf("Error decoding json: %w", err)
	}

	actualConfig = &config{
		next:     areas.Next,
		previous: areas.Previous,
	}

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}
	return nil
}

func commandMapB() error {
	areas := locationAreas{}

	if actualConfig.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := actualConfig.previous

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error in poke api request: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&areas); err != nil {
		return fmt.Errorf("Error decoding json: %w", err)
	}

	actualConfig = &config{
		next:     areas.Next,
		previous: areas.Previous,
	}

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}
	return nil
}
