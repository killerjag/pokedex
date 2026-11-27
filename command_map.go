package main

import (
	"fmt"

	"github.com/killerjag/pokedex/internal/pokeapi"
)

func commandMap(cfg *config, arg string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if cfg.next != nil {
		url = *cfg.next
	}

	location_list := pokeapi.ListLocations(url, cfg.client)

	cfg.next = location_list.Next
	cfg.previous = location_list.Previous

	for _, v := range location_list.Results {
		fmt.Println(v.Name)
	}

	return nil
}
