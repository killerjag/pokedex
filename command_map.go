package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var location_area locationAreaRes
	json.Unmarshal(data, &location_area)

	cfg.next = location_area.Next
	cfg.previous = location_area.Previous

	for _, v := range location_area.Results {
		fmt.Println(v.Name)
	}

	return nil
}
