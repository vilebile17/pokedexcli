package pokeapi

import (
	"fmt"

	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

type Config struct {
	NextLocationsURL string
	PrevLocationsURL string
}

type LocationAreas struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func CommandMap(config *Config, cache *pokecache.Cache, _ *Pokedex, _ string) error {
	var locationAreas LocationAreas
	var err error
	if config.NextLocationsURL == "" {
		locationAreas, err = MakeRequest("https://pokeapi.co/api/v2/location-area/", cache)
	} else {
		locationAreas, err = MakeRequest(config.NextLocationsURL, cache)
	}

	if err != nil {
		return err
	}

	config.NextLocationsURL = locationAreas.Next
	config.PrevLocationsURL = locationAreas.Previous
	return nil
}

func CommandMapb(config *Config, cache *pokecache.Cache, _ *Pokedex, _ string) error {
	if config.PrevLocationsURL == "" {
		fmt.Println("you're on the first page mate :)")
		return nil
	}

	var locationAreas LocationAreas
	var err error
	locationAreas, err = MakeRequest(config.PrevLocationsURL, cache)
	if err != nil {
		return err
	}

	config.NextLocationsURL = locationAreas.Next
	config.PrevLocationsURL = locationAreas.Previous
	return nil
}
