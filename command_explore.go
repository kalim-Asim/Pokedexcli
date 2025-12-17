package main

import (
	"errors"
	"github.com/fatih/color"
)

func callbackExplore(cfg* config, args []string) error {
	if len(args) < 1 {
		return errors.New("please provide a location area name")
	}

	areaName := args[0]
	fullURL := baseURL + "/location-area/" + areaName

	response, err := cfg.pokeapiClient.ListPokemons(fullURL)
	if err != nil {
		return err 
	}

	color.Blue("Exploring %s...\n", areaName)
	color.Blue("Found Pokemons:")
	for _, pokemon := range response.PokemonEncounters {
		color.Magenta(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}