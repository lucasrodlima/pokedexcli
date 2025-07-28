package main

import (
	"bufio"
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/pokeapi"
	"os"
)

func main() {

	for _, c := range cliCommands {
		helpMessage += c.Name + ": " + c.Description + "\n"
	}

	conf := pokeapi.Config{}

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
