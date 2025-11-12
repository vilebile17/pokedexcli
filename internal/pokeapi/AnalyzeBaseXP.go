package pokeapi

// This file is for a test I did to see what base experience each pokemon had in order to get a good probablility for the random catching

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Analyze() int {
	XPtoCount := make(map[int]int)
	i := 1

	for {
		url := "https://pokeapi.co/api/v2/pokemon/" + fmt.Sprint(i)
		var res *http.Response
		res, err := http.Get(url)
		if err != nil {
			break
		}
		if res.StatusCode == 404 {
			break
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			break
		}

		var pokemon Pokemon
		if err = json.Unmarshal(body, &pokemon); err != nil {
			break
		}

		if _, ok := XPtoCount[pokemon.BaseExperience]; ok {
			XPtoCount[pokemon.BaseExperience]++
		} else {
			XPtoCount[pokemon.BaseExperience] = 1
		}

		fmt.Println(i)
		i++
	}

	return WeightedMean(XPtoCount)
}

func WeightedMean(moop map[int]int) int {
	total := 0
	pokemonCount := 0
	for XP := range moop {
		weight := moop[XP]
		total += weight * XP
		pokemonCount += weight
	}

	return total / pokemonCount
}
