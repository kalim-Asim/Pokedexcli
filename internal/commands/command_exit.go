package commands

import (
	"os"
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/shared"
)

func CallbackExit(cfg* shared.Config, args []string) error {
	color.Red("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}