package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/d-shames3/pokedex/internal/pokecache"
)

func main() {
	newScanner := bufio.NewScanner(os.Stdin)
	config := &apiCallConfig{}
	cache := pokecache.NewCache(5 * time.Minute)
	for {
		fmt.Print("Pokedex > ")
		newScanner.Scan()
		userInput := newScanner.Text()
		userInputClean := cleanInput(userInput)
		if len(userInputClean) == 0 {
			fmt.Println("no user input, try again")
			continue
		}
		if len(userInputClean) == 1 {
			userInputClean = append(userInputClean, "")
		}

		cliOptions := getCliCommands()
		cmd, exists := cliOptions[userInputClean[0]]
		if exists {
			err := cmd.callback(cache, config, userInputClean[1])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
