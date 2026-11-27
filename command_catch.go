package main

import (
	"fmt"

	"github.com/killerjag/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, pokemon string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	pokemon_data := pokeapi.GetPokemonData(url, cfg.client)

	if pokemon_data.Name != "" {
		fmt.Println("Throwing a Pokeball at " + pokemon_data.Name + "...")
		fmt.Println(pokemon_data.Name + " was caught!")
		cfg.pokedex[pokemon_data.Name] = pokemon_data
	}

	return nil
}
