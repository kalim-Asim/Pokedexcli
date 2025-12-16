package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	offset = 0
	limit = 20
)

func (c *Client) ListLocationAreas(pageUrl *string) (locationArea, error) {
	endpoint := "/location-area"

	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

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

	return locArea, nil
}