package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/fatih/color"
)

func (c *Client) ListPokemonDetails(pageUrl string) (Pokemon, error) {

	// caching
	dat, ok := c.cache.Get(pageUrl)
	if ok {
		// cache hit
		color.Green("Cache hit...")
		pokemon := Pokemon{}

		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	color.Red("Cache miss...")
	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// add to cache
	c.cache.Add(pageUrl, data)

	return pokemon, nil
}