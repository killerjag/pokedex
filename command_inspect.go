package main

import "fmt"

func commandInspect(cfg *config, pokemon string) error {
	pkm, ok := cfg.pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught " + pokemon)
		return nil
	}

	fmt.Println("Name: " + pkm.Name)
	fmt.Printf("Height: %d\n", pkm.Height)
	fmt.Printf("Weight: %d\n", pkm.Weight)
	fmt.Println("Stats:")
	for _, stat := range pkm.Stats {
		fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pkm.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

	return nil
}
