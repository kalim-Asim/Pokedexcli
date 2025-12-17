package commands

import (
	"errors"
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/shared"
)

// directly fetches pokemon details from map
func CallbackInspect(cfg* shared.Config, args []string) error {
	if len(args) < 1 {
		return errors.New("provide name of pokemon to inspect")
	}

	pokemon := args[0]
	dat, ok := cfg.CaughtPokemons[pokemon]
	if !ok {
		color.Red("you have not caught that pokemon")
		return nil
	}

	color.Magenta("Name: %s\n", dat.Name)
	color.Magenta("Height: %d\n", dat.Height)
	color.Magenta("Weight: %d\n", dat.Weight)
	color.Magenta("Stats:\n")
	for _, stat := range dat.Stats {
		color.Magenta(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	color.Magenta("Types:\n")
	for _, t := range dat.Types {
		color.Magenta(" - %s", t.Type.Name)
	}

	return nil 
}