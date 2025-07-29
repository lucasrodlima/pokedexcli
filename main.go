package main

import (
	"bufio"
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/commands"
	"github.com/lucasrodlima/pokedexcli/internal/pokeapi"
	"github.com/lucasrodlima/pokedexcli/internal/pokecache"
	"os"
	"time"
)

func main() {
	for _, c := range commands.CliCommands {
		commands.HelpMessage += c.Name + ": " + c.Description + "\n"
	}

	mainCache := pokecache.NewCache(5 * time.Second)
	mainPokedex := pokeapi.Pokedex{
		Captured: map[string]pokeapi.Pokemon{},
	}

	conf := pokeapi.Config{
		Cache:   mainCache,
		Pokedex: mainPokedex,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := pokeapi.CleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}

		command, ok := commands.CliCommands[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.Callback(&conf, cleanedInput[1:])
		// no error handling
	}
}
