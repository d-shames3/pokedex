package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/d-shames3/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func commandExit(config *apiCallConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *apiCallConfig) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range getCliCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(config *apiCallConfig) error {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area"
	}
	currentUrl := config.Next

	res, err := pokeapi.CallPokeApi(config.Next)
	if err != nil {
		return err
	}
	data, err := pokeapi.UnmarshalPokeapiResponse(res)
	if err != nil {
		return err
	}

	config.Previous = currentUrl
	if data.Next != "" {
		config.Next = data.Next
	} else {
		config.Next = ""
	}

	results := data.Results
	if len(results) == 0 {
		fmt.Println("No results returned")
		return nil
	}

	for _, location := range results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*apiCallConfig) error
}

type apiCallConfig struct {
	Next     string
	Previous string
}

func getCliCommands() map[string]cliCommand {
	cliCommands := map[string]cliCommand{
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
			description: "Fetches next 20 locations from the PokeApi",
			callback:    commandMap,
		},
	}
	return cliCommands
}
