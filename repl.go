package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/d-shames3/pokedex/internal/pokeapi"
	"github.com/d-shames3/pokedex/internal/pokecache"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func commandExit(cache *pokecache.Cache, config *apiCallConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cache *pokecache.Cache, config *apiCallConfig) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range getCliCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cache *pokecache.Cache, config *apiCallConfig) error {
	if config.Next == "" && config.Previous != "" {
		fmt.Println("You are on the last page of results")
		return nil
	}

	if config.Next == "" && config.Previous == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area"
	}

	cachedData, ok := cache.Get(config.Next)
	data := pokeapi.LocationResponse{}
	if ok {
		err := json.Unmarshal(cachedData, &data)
		if err != nil {
			return fmt.Errorf("error fetching data from cache")
		}
		fmt.Println("Using data from the cache!")
	} else {
		res, err := pokeapi.CallPokeApi(config.Next)
		if err != nil {
			return err
		}
		result, err := pokeapi.UnmarshalPokeapiResponse(res, "location")
		if err != nil {
			return err
		}
		data, ok := result.(pokeapi.LocationResponse)
		if !ok {
			return fmt.Errorf("unexpected response type")
		}
		dataToCache, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("error writing data to cache")
		}
		cache.Add(config.Next, dataToCache)
	}

	if data.Next != "" {
		config.Next = data.Next
	} else {
		config.Next = ""
	}

	if data.Previous != "" {
		config.Previous = data.Previous
	} else {
		config.Previous = ""
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

func commandMapb(cache *pokecache.Cache, config *apiCallConfig) error {
	if config.Previous == "" {
		fmt.Println("You are on the first page of results")
		return nil
	}

	cachedData, ok := cache.Get(config.Previous)
	data := pokeapi.LocationResponse{}
	if ok {
		err := json.Unmarshal(cachedData, &data)
		if err != nil {
			return fmt.Errorf("error fetching data from the cache")
		}
		fmt.Println("Using data from the cache!")
	} else {
		res, err := pokeapi.CallPokeApi(config.Previous)
		if err != nil {
			return err
		}
		result, err := pokeapi.UnmarshalPokeapiResponse(res, "location")
		if err != nil {
			return err
		}
		data, ok := result.(pokeapi.LocationResponse)
		if !ok {
			return fmt.Errorf("unexpected response type")
		}
		dataToCache, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("error writing data to cache")
		}
		cache.Add(config.Previous, dataToCache)
	}

	if data.Previous != "" {
		config.Previous = data.Previous
	} else {
		config.Previous = ""
	}

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
	callback    func(*pokecache.Cache, *apiCallConfig) error
}

type apiCallConfig struct {
	Next     string
	Previous string
}

func getCliCommands() map[string]cliCommand {
	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
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
		"mapb": {
			name:        "mapb",
			description: "Fetches previous 20 locations from the PokeApi",
			callback:    commandMapb,
		},
	}
	return cliCommands
}
