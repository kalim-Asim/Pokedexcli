package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/shared"
)

func (c *Client) ListLocationAreas(pageUrl *string) (shared.LocationArea, error) {
	endpoint := "/location-area"

	fullURL := BaseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

	// caching
	dat, ok := c.Cache.Get(fullURL)
	if ok {
		// cache hit
		color.Green("Cache hit...")
		locArea := shared.LocationArea{}
		err := json.Unmarshal(dat, &locArea)
		if err != nil {
			return shared.LocationArea{}, err
		}

		return locArea, nil
	}

	color.Red("Cache miss...")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return shared.LocationArea{}, err
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return shared.LocationArea{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return shared.LocationArea{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return shared.LocationArea{}, err
	}

	locArea := shared.LocationArea{}
	err = json.Unmarshal(data, &locArea)
	if err != nil {
		return shared.LocationArea{}, err
	}

	// add to cache
	c.Cache.Add(fullURL, data)

	return locArea, nil
}
