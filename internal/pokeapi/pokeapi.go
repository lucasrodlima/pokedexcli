package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/pokecache"
	"io"
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
	Cache    *pokecache.Cache
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

	val, ok := c.Cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error in poke api request: %w", err)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &areas)
		if err != nil {
			return err
		}

		c.Cache.Add(url, data)

	} else {
		err := json.Unmarshal(val, &areas)
		if err != nil {
			return errors.New("Error unmarshaling cached data")
		}
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

	val, ok := c.Cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error in poke api request: %w", err)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &areas)
		if err != nil {
			return err
		}

		c.Cache.Add(url, data)

	} else {
		err := json.Unmarshal(val, &areas)
		if err != nil {
			return errors.New("Error unmarshaling cached data")
		}
	}

	c.Next = areas.Next
	c.Previous = areas.Previous

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}
