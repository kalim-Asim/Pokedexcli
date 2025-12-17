package commands

import (
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

func CallbackPokedex(cfg *pokeapi.Config, args []string) error {
	color.Yellow("Your Pokedex:\n")
	for pokemon, _ := range cfg.CaughtPokemons {
		color.Yellow(" - %s \n", pokemon)
	}
	return nil 
}