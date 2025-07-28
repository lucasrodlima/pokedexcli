package pokeapi

import (
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/pokecache"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  leading and trailing spaces  ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "Some More         tests",
			expected: []string{"some", "more", "tests"},
		},
		{
			input:    "CHARMANDER  Bulbasaur PikaCHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput output length does not match expected")
		}

		for i := range actual {
			expectedWord := c.expected[i]
			actualWord := actual[i]

			if actualWord != expectedWord {
				t.Errorf("cleanInput output word does not match expected")
			}
		}
	}
}

func TestMapLocationAreas(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"count": 1, "next": "https://new.url.com", "previous": "", "results": [{"name": "test-area", "url": "https://test.com"}]}`)
	}))
	defer server.Close()

	config := &Config{
		Next:  server.URL,
		Cache: pokecache.NewCache(5 * time.Second),
	}

	err := MapLocationAreas(config)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if config.Next != "https://new.url.com" {
		t.Errorf("Next URL not updated correctly")
	}
}

func TestMapBLocationAreas(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"count": 1, "next": "", "previous": "https://old.url.com", "results": [{"name": "old-test-area", "url": "https://old-test.com"}]}`)
	}))
	defer server.Close()

	config := &Config{
		Previous: server.URL,
		Cache:    pokecache.NewCache(5 * time.Second),
	}

	err := MapBLocationAreas(config)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if config.Previous != "https://old.url.com" {
		t.Errorf("Previous URL not updated correctly")
	}
}

func TestMapLocationAreasCaching(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		fmt.Fprintln(w, `{"count": 1, "next": "https://new.url.com", "previous": "", "results": [{"name": "test-area", "url": "https://test.com"}]}`)
	}))
	defer server.Close()

	config := &Config{
		Next:  server.URL,
		Cache: pokecache.NewCache(5 * time.Second),
	}

	MapLocationAreas(config)
	MapLocationAreas(config)

	if requestCount > 1 {
		t.Errorf("Expected request to be cached")
	}
}
