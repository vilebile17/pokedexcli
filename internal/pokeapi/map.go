package pokeapi 

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

type Config struct {
	NextLocationsURL string
	PrevLocationsURL string
}

type LocationAreas struct {
	Count int
	Next string
	Previous string
	Results []struct{Name string; Url string}
}

func makeRequest(url string, cache *pokecache.Cache) (LocationAreas, error) {

	var err error
	if body, ok := cache.Get(url); ok {
		var locationAreas LocationAreas
		if err := json.Unmarshal(body, &locationAreas); err == nil {

			for _, location := range locationAreas.Results {
				fmt.Println(location.Name)
			}
			return locationAreas, nil
  	}
		return LocationAreas{}, err
	}

	var res *http.Response
	res, err = http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	var locationAreas LocationAreas
	if err = json.Unmarshal(body, &locationAreas); err != nil {
		return LocationAreas{}, err
  }

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	cache.Add(url, body)
	return locationAreas, nil
}

func CommandMap(config *Config, cache *pokecache.Cache) error {
	var locationAreas LocationAreas
	var err error
	if config.NextLocationsURL == "" {
		locationAreas, err = makeRequest("https://pokeapi.co/api/v2/location-area/", cache)
	} else {
		locationAreas, err = makeRequest(config.NextLocationsURL, cache)
	}

	if err != nil {
		return err
	}

	config.NextLocationsURL = locationAreas.Next
	config.PrevLocationsURL = locationAreas.Previous
	return nil
}

func CommandMapb(config *Config, cache *pokecache.Cache) error {
	if config.PrevLocationsURL == "" {
		fmt.Println("you're on the first page mate :)")
		return nil
	}

	var locationAreas LocationAreas
	var err error
	locationAreas, err = makeRequest(config.PrevLocationsURL, cache)

	if err != nil {
		return err
	}

	config.NextLocationsURL = locationAreas.Next
	config.PrevLocationsURL = locationAreas.Previous
	return nil
}
