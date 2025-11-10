package pokeapi 

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

type Location struct {
	ID int 
	Name string 
	GameIndex int `json:"game_index"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
type PokemonEncounter struct {
	Pokemon Pokemon
}
type Pokemon struct {
	Name string
	Url string
}

func CommandExplore(_ *Config, _ *pokecache.Cache, place string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + place 

	var err error
	var res *http.Response
	res, err = http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode == 404 {
		return fmt.Errorf("Who's... that... pokemon! (cuz I can't find it)")
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
	return nil
}
