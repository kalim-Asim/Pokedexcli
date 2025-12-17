package commands

import (
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/shared"
)

func CallbackHelp(cfg* shared.Config, args []string) error {
	color.Green("Welcome to the Pokedex!")
	color.Green("Usage: ")
	color.Green(" ")
	
	commands := GetCommands()
	for _, cmd := range commands {
		color.Yellow("%s: %s\n", cmd.Name, cmd.Description)
	}

	return nil
}