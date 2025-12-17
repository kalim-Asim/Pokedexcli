package commands

import (
	"github.com/kalim-Asim/pokedexcli/shared"
)

func GetCommands() map[string]shared.CliCommand {
	return map[string]shared.CliCommand{
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