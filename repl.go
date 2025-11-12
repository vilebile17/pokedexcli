package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vilebile17/pokedexcli/internal/pokeapi"
	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

// We begin by creating a registery of all of the commands

type PokedexCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, *pokecache.Cache, *pokeapi.Pokedex, string) error
}

func makeCommandMap() map[string]PokedexCommand {
	moop := map[string]PokedexCommand{
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
			description: "Prints 20 locations in the Pokemon universe",
			callback:    pokeapi.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous page of Pokemon locations",
			callback:    pokeapi.CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Prints all of the pokemon available in a region. USAGE: explore <location>",
			callback:    pokeapi.CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon!",
			callback:    pokeapi.CommandCatch,
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
	pokedex := &pokeapi.Pokedex{}

	for {
		fmt.Print("Pokedex > ")

		if inputExists := s.Scan(); inputExists {
			input := s.Text()
			cleanedInput := CleanInput(input)
			cleanedInput = append(cleanedInput, "")

			if len(cleanedInput) == 0 { // empty line case
				continue
			}

			command := cleanedInput[0]
			param := cleanedInput[1]

			if _, ok := commandMap[command]; ok {
				err := commandMap[command].callback(mapConfig, cache, pokedex, param)
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

func commandExit(_ *pokeapi.Config, _ *pokecache.Cache, _ *pokeapi.Pokedex, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeapi.Config, _ *pokecache.Cache, _ *pokeapi.Pokedex, param string) error {
	moop := makeCommandMap()

	// this section is for the help message of a single command
	if param != "" {
		if _, ok := moop[param]; ok {
			fmt.Printf("%v: %v\n", moop[param].name, moop[param].description)
		} else {
			fmt.Println("Command not found")
		}
		return nil
	}

	// The remaining part is for the overall help message
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for command := range moop {
		fmt.Printf("%v: %v\n", moop[command].name, moop[command].description)
	}
	fmt.Println("")

	return nil
}
