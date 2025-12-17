package commands

import (
	"errors"
	"math/rand"
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

func CallbackCatch(cfg* pokeapi.Config, args []string) error {
	if len(args) < 1 {
		return errors.New("please provide a pokemon to catch")
	}

	pokemon := args[0]
	fullURL := pokeapi.BaseURL + "/pokemon/" + pokemon 

	response, err := cfg.PokeApiClient.ListPokemonDetails(fullURL)
	if err != nil {
		return err 
	}

	color.Blue("Throwing a Pokeball at %s...\n", pokemon)
	
	// check if already caught
	if _, ok := cfg.CaughtPokemons[pokemon]; ok {
		color.Green("%s was caught!\n", pokemon)
		color.Green("You may now inspect it with the inspect command.")
		return nil 
	}

	randomNumber := rand.Intn(2 * response.BaseExperience)
	if randomNumber >= response.BaseExperience { // can catch
		color.Green("%s was caught!\n", pokemon)
		color.Green("You may now inspect it with the inspect command.")
		cfg.CaughtPokemons[pokemon] = response 
	} else { // cannot catch
		color.Red("%s excaped!\n", pokemon)
	}

	return nil
}