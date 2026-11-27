package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/killerjag/pokedex/internal/pokecache"
)

//const baseURL string = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func ListLocations(url string, c Client) locationAreaRes {
	if val, ok := c.cache.Get(url); ok {
		location_area := locationAreaRes{}
		err := json.Unmarshal(val, &location_area)
		if err != nil {
			return locationAreaRes{}
		}
		return location_area
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaRes{}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaRes{}
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaRes{}
	}

	var location_area locationAreaRes
	json.Unmarshal(data, &location_area)

	c.cache.Add(url, data)

	return location_area
}

func ListPokemon(url string, c Client) exploreAreaRes {
	if val, ok := c.cache.Get(url); ok {
		area_data := exploreAreaRes{}
		err := json.Unmarshal(val, &area_data)
		if err != nil {
			return exploreAreaRes{}
		}
		return area_data
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return exploreAreaRes{}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return exploreAreaRes{}
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return exploreAreaRes{}
	}

	var area_data exploreAreaRes
	json.Unmarshal(data, &area_data)

	c.cache.Add(url, data)

	return area_data
}

func GetPokemonData(url string, c Client) Pokemon {
	return Pokemon{}
}
