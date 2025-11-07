package main

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

type LocationAreas struct {
	Count int
	Next string
	Previous string
	Results []struct{Name string; Url string}
}

func CommandMap(config *Config) error {
	var res *http.Response
	var err error
	if config.nextLocationsURL == "" {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
	} else {
		res, err = http.Get(config.nextLocationsURL)
	}

	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationAreas LocationAreas
	if err = json.Unmarshal(body, &locationAreas); err != nil {
		return err
  }

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationsURL = locationAreas.Next
	config.prevLocationsURL = locationAreas.Previous
	return nil
}

func CommandMapb(config *Config) error {
	if config.prevLocationsURL == "" {
		fmt.Println("you're on the first page mate :)")
		return nil
	}

	res, err := http.Get(config.prevLocationsURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationAreas LocationAreas
	if err = json.Unmarshal(body, &locationAreas); err != nil {
		return err
  }

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationsURL = locationAreas.Next
	config.prevLocationsURL = locationAreas.Previous
	return nil
}
