package main

import (
	"bytes"
	"fmt"
	"github.com/lucasrodlima/pokedexcli/internal/pokeapi"
	"github.com/lucasrodlima/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCommandHelp(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	commandHelp(nil)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if !strings.Contains(buf.String(), "Welcome to the Pokedex!") {
		t.Errorf("Expected help message to be printed")
	}
}

func TestCommandMap(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"count": 1, "next": "https://new.url.com", "previous": "", "results": [{"name": "test-area", "url": "https://test.com"}]}`)
	}))
	defer server.Close()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	config := &pokeapi.Config{
		Next:  server.URL,
		Cache: pokecache.NewCache(5 * time.Second),
	}
	commandMap(config)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if !strings.Contains(buf.String(), "test-area") {
		t.Errorf("Expected location area to be printed")
	}
}

func TestCommandMapB(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"count": 1, "next": "", "previous": "https://old.url.com", "results": [{"name": "old-test-area", "url": "https://old-test.com"}]}`)
	}))
	defer server.Close()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	config := &pokeapi.Config{
		Previous: server.URL,
		Cache:    pokecache.NewCache(5 * time.Second),
	}
	commandMapB(config)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if !strings.Contains(buf.String(), "old-test-area") {
		t.Errorf("Expected location area to be printed")
	}
}
