package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"github.com/vilebile17/pokedexcli/internal/pokecache"
)

func MakeRequest(url string, cache *pokecache.Cache) (LocationAreas, error) {

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
