package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

type Pokedex map[string]Pokemon

func CommandCatch(_ *Config, c *pokecache.Cache, p *Pokedex, param string) error {
	var err error
	if body, ok := c.Get(param); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(body, &pokemon); err == nil {
			return CatchOrNot(pokemon, p)
		}
		return err
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + param
	var res *http.Response
	res, err = http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode == 404 {
		return fmt.Errorf("Who's... that... Pokemon! Cause I don't know who it is\n")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	if err = json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	c.Add(param, body)
	return CatchOrNot(pokemon, p)
}

func CatchOrNot(pokemon Pokemon, pokedex *Pokedex) error {
	fmt.Printf("Throwing a Pokeball at %v\n", pokemon.Name)

	num := rand.Intn(500)
	if num < pokemon.BaseExperience {
		// consider the base experience to be the pokemon's strength; if it is stronger it is more likely to resist a pokeball
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%v was caught!\n", pokemon.Name)
	(*pokedex)[pokemon.Name] = pokemon
	return nil
}
