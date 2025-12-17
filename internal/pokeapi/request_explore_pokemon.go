package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/kalim-Asim/pokedexcli/shared"
	"github.com/fatih/color"
)

func (c* Client) ListPokemons(pageUrl string) (shared.ExploreLocation, error) {

	// caching 
	dat, ok := c.Cache.Get(pageUrl)
	if ok {
		// cache hit 
		color.Green("Cache hit...")
		pokemon := shared.ExploreLocation{}

		err := json.Unmarshal(dat, &pokemon) 
		if err != nil {
			return shared.ExploreLocation{}, err  
		}

		return pokemon, nil 
	}

	color.Red("Cache miss...")
	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return shared.ExploreLocation{}, err 
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return shared.ExploreLocation{}, err 
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return shared.ExploreLocation{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return shared.ExploreLocation{}, err 
	}

	pokemons := shared.ExploreLocation{}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return shared.ExploreLocation{}, err 
	}

	// add to cache
	c.Cache.Add(pageUrl, data)

	return pokemons, nil
}