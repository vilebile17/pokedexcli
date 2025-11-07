package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
) 

// We begin my creating a registery of all of the commands
type pokedexCommand struct {
	name string
	description string
	callback func() error
}
func makeCommandMap() map[string]pokedexCommand {
	moop := map[string]pokedexCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	} 
	return moop
}

// Takes each word in the standard input and puts them in a slice
func cleanInput(text string) []string {
	newText := strings.ToLower(text)
	return strings.Fields(newText)	
}

// main command
func StartRepl() {
	s := bufio.NewScanner(os.Stdin)
	commandMap := makeCommandMap()

	for true {
		fmt.Print("Pokedex > ")

		if inputExists := s.Scan(); inputExists {
			input := s.Text()
			cleanedInput := cleanInput(input)
			command := cleanedInput[0]

			if _, ok := commandMap[command]; ok {
				commandMap[command].callback()
			} else {
				fmt.Println("Unknown command")
			}

		}
	}
}

// Here begin all of the functions for the each of the commands

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	moop := makeCommandMap()	
	for command := range moop {
		fmt.Printf("%v: %v\n", moop[command].name, moop[command].description)
	}
	fmt.Println("")

	return nil
}
