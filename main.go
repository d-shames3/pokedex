package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		newScanner.Scan()
		userInput := newScanner.Text()
		userInputClean := cleanInput(userInput)
		if len(userInputClean) == 0 {
			fmt.Println("no user input, try again")
			continue
		}
		cliOptions := getCliCommands()
		cmd, exists := cliOptions[userInputClean[0]]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
