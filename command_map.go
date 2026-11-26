package main

import (
	"fmt"

	"github.com/killerjag/pokedex/internal/pokeapi"
)

type locationAreaRes struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

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
