package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg* config) error {
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, areas := range response.Results {
		fmt.Printf(" - %s\n", areas.Name)
	}

	cfg.nextLocationUrl = response.Next
	cfg.prevLocationUrl = response.Previous
	return nil
}