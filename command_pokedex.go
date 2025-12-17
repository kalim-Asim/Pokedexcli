package main

import "github.com/fatih/color"

func callbackPokedex(cfg *config, args []string) error {
	color.Yellow("Your Pokedex:\n")
	for pokemon, _ := range cfg.caughtPokemons {
		color.Yellow(" - %s \n", pokemon)
	}
	return nil 
}