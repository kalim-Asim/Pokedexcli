package main

import (
	"github.com/fatih/color"
)

func callbackHelp(cfg* config, args []string) error {
	color.Green("Welcome to the Pokedex!")
	color.Green("Usage: ")
	color.Green(" ")

	commands := getCommands()
	for _, cmd := range commands {
		color.Yellow("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}