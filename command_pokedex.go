package main

import "fmt"

func commandPokedex(cfg *config, arg string) error {
	fmt.Println("Your Pokedex:")
	for _, pkm := range cfg.pokedex {
		fmt.Printf(" - %s\n", pkm.Name)
	}

	return nil
}
