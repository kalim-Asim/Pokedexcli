package commands

import (
	"os"

	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

func CallbackExit(cfg* pokeapi.Config, args []string) error {
	color.Red("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}