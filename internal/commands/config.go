package commands

import (
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config, []string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"pokedex": {
			Name: 				"pokedex",
			Description:  "Lists all caught pokemons",
			Callback: 		CallbackPokedex,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CallbackHelp,
		},
		"map": {
			Name:        "map",
			Description: "Lists next 20 location",
			Callback:    CallbackMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Lists previous 20 location",
			Callback:    CallbackMapb,
		},
		"explore": {
			Name:					"explore",
			Description: 	"Explore any particular location",
			Callback: 		CallbackExplore,
		},
		"catch": {
			Name: 			"catch",
			Description: "Catch a particular pokemon",
			Callback:    CallbackCatch,
		},
		"inspect": {
			Name: 				"inspect",
			Description: 	"Details about caught pokemon",
			Callback: 		CallbackInspect,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CallbackExit,
		},
	}
}