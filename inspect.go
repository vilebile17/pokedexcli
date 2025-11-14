package main

import (
	"fmt"
	"strings"

	"github.com/vilebile17/pokedexcli/internal/pokeapi"
	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

func CommandInspect(_ *pokeapi.Config, _ *pokecache.Cache, p *pokeapi.Pokedex, param string) error {
	details, ok := (*p)[param]
	if !ok {
		fmt.Println("You don't seem to have caught that pokemon yet...")
		return nil
	}

	fmt.Print(strings.ToTitle(fmt.Sprintf("===== %v =====\n", details.Name)))
	fmt.Printf("ID: %v\n", details.ID)
	fmt.Printf("Height: %.1fm\n", float64(details.Height)*0.1)
	fmt.Printf("Weight: %.1fkg\n", float64(details.Weight)*0.1)
	fmt.Println("Types:")
	for _, Type := range details.Types {
		fmt.Printf("   - %v\n", Type.Type.Name)
	}
	fmt.Println("Stats:")
	for _, s := range details.Stats {
		fmt.Printf("   %v - %v\n", s.Stat.Name, s.BaseStat)
	}

	return nil
}
