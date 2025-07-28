package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type locationAreas struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {

	for _, c := range cliCommands {
		helpMessage += c.name + ": " + c.description + "\n"
	}

	conf := config{}

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
		command.callback(&conf)
	}
}
