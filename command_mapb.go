package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMapb(cfg* config) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("there is no previous Url")
	}
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas: ")
	for _, areas := range response.Results {
		fmt.Println(areas.Name)
	}
  
	cfg.nextLocationUrl = response.Next
	cfg.prevLocationUrl = response.Previous
	return nil
}