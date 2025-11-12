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
	fmt.Printf("Height: %v\n", details.Height)
	fmt.Printf("Weight: %v\n", details.Weight)
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
