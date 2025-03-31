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
		} else {
			fmt.Printf("Your command was: %s\n", userInputClean[0])
		}
	}

}
