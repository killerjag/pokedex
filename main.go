package main

import (
	"time"

	"github.com/killerjag/pokedex/internal/pokeapi"
)

func main() {
	pokeclient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		client:  pokeclient,
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
