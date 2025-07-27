package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
