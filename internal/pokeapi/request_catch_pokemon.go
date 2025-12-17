package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/shared"
)

func (c *Client) ListPokemonDetails(pageUrl string) (shared.Pokemon, error) {
	// caching
	dat, ok := c.Cache.Get(pageUrl)
	if ok {
		// cache hit
		color.Green("Cache hit...")
		pokemon := shared.Pokemon{}

		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return shared.Pokemon{}, err
		}

		return pokemon, nil
	}

	color.Red("Cache miss...")
	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return shared.Pokemon{}, err
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return shared.Pokemon{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return shared.Pokemon{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return shared.Pokemon{}, err
	}

	pokemon := shared.Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return shared.Pokemon{}, err
	}

	// add to cache
	c.Cache.Add(pageUrl, data)

	return pokemon, nil
}
