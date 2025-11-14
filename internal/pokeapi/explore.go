package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

type Location struct {
	ID                int
	Name              string
	GameIndex         int                `json:"game_index"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
type PokemonEncounter struct {
	Pokemon Pokemon
}
type Pokemon struct {
	ID             int
	Name           string
	Height         int
	Weight         int
	BaseExperience int `json:"base_experience"`
	Types          []struct {
		Slot int
		Type struct {
			Name string
			URL  string
		}
	}
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int
		Stat     struct {
			Name string
			URL  string
		}
	}
}

func CommandExplore(_ *Config, c *pokecache.Cache, _ *Pokedex, param string) error {
	var err error
	if body, ok := c.Get(param); ok {
		var location Location
		if err := json.Unmarshal(body, &location); err == nil {

			for _, encounter := range location.PokemonEncounters {
				fmt.Println(encounter.Pokemon.Name)
			}
			return nil
		}
		return err
	}

	url := "https://pokeapi.co/api/v2/location-area/" + param
	var res *http.Response
	res, err = http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode == 404 {
		return fmt.Errorf("Err, I don't know where that is...\n")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var location Location
	if err = json.Unmarshal(body, &location); err != nil {
		return err
	}

	for _, pokemon := range location.PokemonEncounters {
		realPokemon := pokemon.Pokemon // that's a confusing line
		fmt.Println(realPokemon.Name)
	}

	c.Add(param, body)
	return nil
}
