package commands

import (
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/pokeapi"
	"os"
)

var CliCommands = map[string]pokeapi.CliCommand{
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Callback:    CommandHelp,
	},
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
	"map": {
		Name:        "map",
		Description: "List Location Areas",
		Callback:    CommandMap,
	},
	"mapb": {
		Name:        "mapb",
		Description: "List Previous Location Areas",
		Callback:    CommandMapB,
	},
	"explore": {
		Name:        "explore",
		Description: "Explore the Pokemon in area",
		Callback:    CommandExplore,
	},
	"catch": {
		Name:        "catch",
		Description: "Catch a Pokemon and add it to your Pokedex",
		Callback:    commandCatch,
	},
}

var HelpMessage string

func commandCatch(c *pokeapi.Config, args []string) error {
	err := pokeapi.CatchPokemon(c, args[0])
	if err != nil {
		return err
	}

	return nil
}

func CommandHelp(c *pokeapi.Config, args []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	fmt.Println(HelpMessage)
	return nil
}

func CommandExit(c *pokeapi.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandExplore(c *pokeapi.Config, args []string) error {
	fmt.Printf("Exploring %s...\n", args[0])
	err := pokeapi.ExploreArea(c, args[0])
	if err != nil {
		return err
	}

	return nil
}

func CommandMap(c *pokeapi.Config, args []string) error {
	err := pokeapi.MapLocationAreas(c)
	if err != nil {
		return err
	}

	return nil
}

func CommandMapB(c *pokeapi.Config, args []string) error {
	err := pokeapi.MapBLocationAreas(c)
	if err != nil {
		return err
	}

	return nil
}
