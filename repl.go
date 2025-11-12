package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
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
			description: "Attemps to catch a pokemon. USAGE: catch <pokemon>",
			callback:    pokeapi.CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Take a look at a pokemon's details. USAGE: inspect <pokemon>",
			callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all of the pokemon that you have caught",
			callback:    commandPokedex,
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

	// mean := pokeapi.Analyze()
	// fmt.Println(mean)

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

func commandPokedex(_ *pokeapi.Config, _ *pokecache.Cache, p *pokeapi.Pokedex, _ string) error {
	if len(*p) == 0 {
		fmt.Println("You've not got any pokemon yet...")
		return nil
	}

	fmt.Println("Your pokedex:")
	for _, pokemon := range *p {
		Type := pokemon.Types[0].Type.Name
		switch Type {
		case "fire":
			color.Red("   %v\n", pokemon.Name)
		case "grass":
			color.Green("   %v\n", pokemon.Name)
		case "water":
			color.Blue("   %v\n", pokemon.Name)
		case "fighting":
			color.RGB(255, 165, 0).Printf("   %v\n", pokemon.Name)
		case "poison":
			color.RGB(128, 0, 128).Printf("   %v\n", pokemon.Name)
		case "electric":
			color.Yellow("   %v\n", pokemon.Name)
		case "psychic":
			color.Magenta("   %v\n", pokemon.Name)
		case "ice":
			color.Cyan("   %v\n", pokemon.Name)
		default:
			fmt.Printf("   %v\n", pokemon.Name)
		}

	}
	return nil
}
