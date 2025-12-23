package pokeapi

import (
	"encoding/json"
	"errors"
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
		return fmt.Errorf("Who's... that... Pokemon! Cause I don't know who it is")
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

	terapagos(&pokemon)
	return catch(pokemon, p)
}

func catch(pokemon Pokemon, pokedex *Pokedex) error {
	prob := sigmoid(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %v...(%.0f percent)\n", pokemon.Name, prob*100.0)

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

func terapagos(p *Pokemon) {
	if p.Name != "terapagos" {
		return
	}
	p.Types = append(p.Types, p.Types[0])

	p.Types[0] = struct {
		Slot int
		Type struct {
			Name string
			URL  string
		}
	}{
		Slot: 1,
		Type: struct {
			Name string
			URL  string
		}{
			Name: "tera",
			URL:  "https://www.youtube.com/watch?v=At8v_Yc044Y",
		},
	}
}

func CommandFree(_ *Config, _ *pokecache.Cache, p *Pokedex, param string) error {
	if _, ok := (*p)[param]; !ok {
		return errors.New("You don't have that pokemon")
	}

	delete(*p, param)
	fmt.Printf("Goodbye %v, it was nice knowing you!\n", param)
	return nil
}
