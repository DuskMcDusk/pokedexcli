package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/DuskMcDusk/pokedexcli/internal/pokeapi"
)

type config struct {
	prevLocationURL *string
	nextLocationURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func GetCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets the locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous locations",
			callback:    commandMapb,
		},
	}
}

func cleaninput(text string) []string {
	lower := strings.ToLower(text)
	entries := strings.Fields(lower)
	return entries
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for key, command := range GetCommandMap() {
		fmt.Printf("%v: %v\n", key, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationResult, err := pokeapi.GetPokeMap(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = locationResult.Next
	cfg.prevLocationURL = locationResult.Previous

	for _, location := range locationResult.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}
	locationResult, err := pokeapi.GetPokeMap(cfg.prevLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = locationResult.Next
	cfg.prevLocationURL = locationResult.Previous

	for _, location := range locationResult.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleaninput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := GetCommandMap()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}
