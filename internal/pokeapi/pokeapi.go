package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func CatchPokemon(c *Config, name string) error {
	pokemon := Pokemon{}

	url := "https://pokeapi.co/api/v2/pokemon/" + name

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

		err = json.Unmarshal(data, &pokemon)
		if err != nil {
			return err
		}

		c.Cache.Add(url, data)

	} else {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return errors.New("Error unmarshaling cached data")
		}
	}

	_, ok = c.Pokedex.Captured[pokemon.Name]
	if !ok {
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
		time.Sleep(time.Second)
		if rand.Intn(pokemon.BaseExperience) < 50 {
			fmt.Printf("%s was caught!\nYou may inspect it with the inspect command.\n", pokemon.Name)
			c.Pokedex.Captured[pokemon.Name] = pokemon
		} else {
			fmt.Printf("%s escaped!\n", pokemon.Name)
		}
	} else {
		fmt.Printf("%s is already caught!\n", pokemon.Name)
	}

	return nil
}

func ExploreArea(c *Config, area string) error {
	var explored exploredArea

	url := "https://pokeapi.co/api/v2/location-area/" + area

	fmt.Println("Found Pokemon:")

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

		err = json.Unmarshal(data, &explored)
		if err != nil {
			return err
		}

		c.Cache.Add(url, data)

	} else {
		err := json.Unmarshal(val, &explored)
		if err != nil {
			return errors.New("Error unmarshaling cached data")
		}
	}

	for _, pe := range explored.PokemonEncounters {
		pokemonName := pe.Pokemon.Name
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
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
