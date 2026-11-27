package main

import "github.com/killerjag/pokedex/internal/pokeapi"

func commandCatch(cfg *config, pokemon string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	pokemon_data := pokeapi.GetPokemonData(url, cfg.client)

	println("Throwing a Pokeball at " + pokemon_data.Name)

	return nil
}
