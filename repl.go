package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
) 

// We begin by creating a registery of all of the commands
type Config struct {
	nextLocationsURL string
	prevLocationsURL string
}
type PokedexCommand struct {
	name string
	description string
	callback func(*Config) error
}
func makeCommandMap() map[string]PokedexCommand {
	moop := map[string]PokedexCommand{
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
		"map": {
			name: "map",
			description: "Prints 20 locations in the Pokemon universe",
			callback: CommandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Prints the previous page of Pokemon locations",
			callback: CommandMapb,
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
	mapConfig := Config{nextLocationsURL: "", prevLocationsURL: ""}

	for true {
		fmt.Print("Pokedex > ")

		if inputExists := s.Scan(); inputExists {
			input := s.Text()
			cleanedInput := cleanInput(input)

			if len(cleanedInput) == 0 { // empty line case
				continue
			}

			command := cleanedInput[0]

			if _, ok := commandMap[command]; ok {
				commandMap[command].callback(&mapConfig)
			} else {
				fmt.Println("Unknown command")
			}

		}
	}
}

// Here begin all of the functions for the each of the commands

func commandExit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(_ *Config) error {
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
