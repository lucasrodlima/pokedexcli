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

type exploredArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
	Callback    func(*Config, []string) error
	Args        []string
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
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
