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
	"inspect": {
		Name:        "inspect",
		Description: "Inspect information about caught Pokemon",
		Callback:    commandInspect,
	},
	"pokedex": {
		Name:        "pokedex",
		Description: "List all caught Pokemon",
		Callback:    commandPokedex,
	},
}

func commandPokedex(c *pokeapi.Config, args []string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range c.Pokedex.Captured {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}

func commandInspect(c *pokeapi.Config, args []string) error {
	pokemon, ok := c.Pokedex.Captured[args[0]]
	if !ok {
		fmt.Printf("%s is not caught!\n", args[0])
		return nil
	}

	fmt.Printf(`Name: %v
Height: %v
Weight: %v
Stats:
  -hp: %v
  -attack: %v
  -defense: %v
  -special-attack: %v
  -special-defense: %v
  -speed: %v
Types:
`, pokemon.Name, pokemon.Height, pokemon.Weight,
		pokemon.Stats[0].BaseStat, pokemon.Stats[1].BaseStat,
		pokemon.Stats[2].BaseStat, pokemon.Stats[3].BaseStat,
		pokemon.Stats[4].BaseStat, pokemon.Stats[5].BaseStat,
	)

	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func commandCatch(c *pokeapi.Config, args []string) error {
	err := pokeapi.CatchPokemon(c, args[0])
	if err != nil {
		return err
	}

	return nil
}

var HelpMessage string

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
