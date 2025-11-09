package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/vilebile17/pokedexcli/internal/pokeapi"
	"github.com/vilebile17/pokedexcli/internal/pokecache"
) 

// We begin by creating a registery of all of the commands

type PokedexCommand struct {
	name string
	description string
	callback func(*pokeapi.Config, *pokecache.Cache) error
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
			callback: pokeapi.CommandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Prints the previous page of Pokemon locations",
			callback: pokeapi.CommandMapb,
		},
	} 
	return moop
}

// Takes each word in the standard input and puts them in a slice
func CleanInput(text string) []string {
	newText := strings.ToLower(text)
	return strings.Fields(newText)	
}

// main command
func StartRepl() {
	s := bufio.NewScanner(os.Stdin)
	commandMap := makeCommandMap()

	mapConfig := &(pokeapi.Config{NextLocationsURL: "", PrevLocationsURL: ""})
	cache := pokecache.NewCache(5 * time.Second)

	for true {
		fmt.Print("Pokedex > ")

		if inputExists := s.Scan(); inputExists {
			input := s.Text()
			cleanedInput := CleanInput(input)

			if len(cleanedInput) == 0 { // empty line case
				continue
			}

			command := cleanedInput[0]

			if _, ok := commandMap[command]; ok {
				err := commandMap[command].callback(mapConfig, cache)
				if err != nil {
					fmt.Print(err)
				}
			} else {
				fmt.Println("Unknown command - try running 'help' if you're stuck")
			}

		}
	}
}

// Here begin all of the functions for the each of the commands

func commandExit(_ *pokeapi.Config, _ *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(_ *pokeapi.Config, _ *pokecache.Cache) error {
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
