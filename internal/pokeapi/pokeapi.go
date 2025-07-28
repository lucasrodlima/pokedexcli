package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type LocationAreas struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

type Config struct {
	Next     string
	Previous string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func MapLocationAreas(c *Config) error {
	areas := LocationAreas{}
	var url string

	if c.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = c.Next
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

	c.Next = areas.Next
	c.Previous = areas.Previous

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func MapBLocationAreas(c *Config) error {
	areas := LocationAreas{}

	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := c.Previous

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error in poke api request: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&areas); err != nil {
		return fmt.Errorf("Error decoding json: %w", err)
	}

	c.Next = areas.Next
	c.Previous = areas.Previous

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}
