package main

import (
	"fmt"
)

func commandHelp(cfg *config, arg string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
