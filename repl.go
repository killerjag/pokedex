package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/killerjag/pokedex/internal/pokeapi"
)

type config struct {
	client   pokeapi.Client
	next     *string
	previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, arg string) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			continue
		}
		input := scanner.Text()
		cleaned_input := cleanInput(input)
		command := cleaned_input[0]
		arg := ""
		if len(cleaned_input) > 1 {
			arg = cleaned_input[1]
		}
		cmd, ok := commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback(cfg, arg)
	}
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch the given pokemon",
			callback:    commandCatch,
		},
	}
	return commands
}
