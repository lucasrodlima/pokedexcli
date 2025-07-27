package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	cliCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			usageLines := ""
			for _, c := range cliCommands {
				usageLines += c.name + ": " + c.description + "\n"
			}
			fmt.Printf(`Welcome to the Pokedex!
Usage:

%s
`, usageLines)
			return nil
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}

		command, ok := cliCommands[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}
}
