package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
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
			return catch(pokemon, p)
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
	return catch(pokemon, p)
}

func catch(pokemon Pokemon, pokedex *Pokedex) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	prob := sigmoid(pokemon.BaseExperience)
	fmt.Printf("The chance that %v is caught is %.2f percent\n", pokemon.Name, prob*100.0)

	num := rand.Float64()
	if num < prob {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		(*pokedex)[pokemon.Name] = pokemon
		return nil
	}

	fmt.Printf("%v escaped!\n", pokemon.Name)
	return nil
}

func sigmoid(x int) float64 {
	power := 0.01 * (float64(x) - 150.0)
	quotient := 1.0 + math.Pow(math.E, power)
	return 1.0 / quotient
}
