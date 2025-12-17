package main

import (
	"errors"
	"github.com/fatih/color"
)

// directly fetches pokemon details from map
func callbackInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("provide name of pokemon to inspect")
	}

	pokemon := args[0]
	dat, ok := cfg.caughtPokemons[pokemon]
	if !ok {
		return errors.New("you have not caught that pokemon")
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