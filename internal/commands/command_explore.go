package commands

import (
	"errors"
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

func CallbackExplore(cfg* pokeapi.Config, args []string) error {
	if len(args) < 1 {
		return errors.New("please provide a location area name")
	}

	areaName := args[0]
	fullURL := pokeapi.BaseURL + "/location-area/" + areaName

	response, err := cfg.PokeApiClient.ListPokemons(fullURL)
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