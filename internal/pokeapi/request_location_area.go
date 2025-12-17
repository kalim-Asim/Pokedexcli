package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

func (c *Client) ListLocationAreas(pageUrl *string) (locationArea, error) {
	endpoint := "/location-area"

	fullURL := BaseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

	// caching
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit 
		color.Green("Cache hit...")
		locArea := locationArea{}
		err := json.Unmarshal(dat, &locArea)
		if err != nil {
			return locationArea{}, err 
		}

		return locArea, nil
	}

	color.Red("Cache miss...")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationArea{}, err 
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err 
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return locationArea{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return locationArea{}, err 
	}

	locArea := locationArea{}
	err = json.Unmarshal(data, &locArea)
	if err != nil {
		return locationArea{}, err 
	}

	// add to cache
	c.cache.Add(fullURL, data)

	return locArea, nil
}