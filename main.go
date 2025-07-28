package main

import (
	"bufio"
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/pokeapi"
	"github.com/lucasrodlima/pokedexcli/internal/pokecache"
	"os"
	"time"
)

func main() {
	for _, c := range cliCommands {
		helpMessage += c.Name + ": " + c.Description + "\n"
	}

	mainCache := pokecache.NewCache(5 * time.Second)

	conf := pokeapi.Config{
		Cache: mainCache,
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

		command, ok := cliCommands[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.Callback(&conf)
	}
}
