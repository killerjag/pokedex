package main

import (
	"fmt"

	"github.com/killerjag/pokedex/internal/pokeapi"
)

func commandMapb(cfg *config, arg string) error {
	if cfg.previous == nil {
		fmt.Println("Can't go back")
		return nil
	}

	url := *cfg.previous

	location_list := pokeapi.ListLocations(url, cfg.client)

	cfg.next = location_list.Next
	cfg.previous = location_list.Previous

	for _, v := range location_list.Results {
		fmt.Println(v.Name)
	}

	return nil
}
