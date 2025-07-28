package main

import (
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/pokeapi"
	"os"
)

var cliCommands = map[string]pokeapi.CliCommand{
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commandHelp,
	},
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	},
	"map": {
		Name:        "map",
		Description: "List Location Areas",
		Callback:    commandMap,
	},
	"mapb": {
		Name:        "mapb",
		Description: "List Previous Location Areas",
		Callback:    commandMapB,
	},
}

var helpMessage string

func commandHelp(c *pokeapi.Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	fmt.Println(helpMessage)
	return nil
}

func commandExit(c *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *pokeapi.Config) error {
	err := pokeapi.MapLocationAreas(c)
	if err != nil {
		return err
	}

	return nil
}

func commandMapB(c *pokeapi.Config) error {
	err := pokeapi.MapBLocationAreas(c)
	if err != nil {
		return err
	}

	return nil
}
