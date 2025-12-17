package main

import (
	"errors"
	"github.com/fatih/color"
)

func callbackMap(cfg* config, args []string) error {
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	for _, areas := range response.Results {
		color.Magenta(" - %s\n", areas.Name)

	}

	cfg.nextLocationUrl = response.Next
	cfg.prevLocationUrl = response.Previous
	return nil
}

func callbackMapb(cfg* config, args []string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("there is no previous Url")
	}

	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	for _, areas := range response.Results {
		color.Magenta(" - %s\n", areas.Name)
	}
  
	cfg.nextLocationUrl = response.Next
	cfg.prevLocationUrl = response.Previous
	return nil
}