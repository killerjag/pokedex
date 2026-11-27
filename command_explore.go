package main

import (
	"fmt"

	"github.com/killerjag/pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, location string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + location

	area_data := pokeapi.ListPokemon(url, cfg.client)

	println("Exploring " + area_data.Name + "...")
	println("Found Pokemon:")

	for _, v := range area_data.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}

	return nil
}
