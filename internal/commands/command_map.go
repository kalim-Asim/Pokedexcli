package commands

import (
	"errors"
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/shared"
)

func CallbackMap(cfg* shared.Config, args []string) error {
	response, err := cfg.PokeApiClient.ListLocationAreas(cfg.NextLocationUrl)
	if err != nil {
		return err
	}

	for _, areas := range response.Results {
		color.Magenta(" - %s\n", areas.Name)
	}

	cfg.NextLocationUrl = response.Next
	cfg.PrevLocationUrl = response.Previous
	return nil
}

func CallbackMapb(cfg* shared.Config, args []string) error {
	if cfg.PrevLocationUrl == nil {
		return errors.New("there is no previous Url")
	}

	response, err := cfg.PokeApiClient.ListLocationAreas(cfg.PrevLocationUrl)
	if err != nil {
		return err
	}

	for _, areas := range response.Results {
		color.Magenta(" - %s\n", areas.Name)
	}
  
	cfg.NextLocationUrl = response.Next
	cfg.PrevLocationUrl = response.Previous
	return nil
}