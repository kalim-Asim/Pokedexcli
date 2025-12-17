package main

import (
	"errors"
	"math/rand"
	"github.com/fatih/color"
)

func callbackCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("please provide a pokemon to catch")
	}

	pokemon := args[0]
	fullURL := baseURL + "/pokemon/" + pokemon 

	response, err := cfg.pokeapiClient.ListPokemonDetails(fullURL)
	if err != nil {
		return err 
	}

	color.Blue("Throwing a Pokeball at %s...\n", pokemon)
	
	// check if already caught
	if _, ok := cfg.caughtPokemons[pokemon]; ok {
		color.Green("%s was caught!\n", pokemon)
		return nil 
	}

	randomNumber := rand.Intn(2 * response.BaseExperience)
	if randomNumber >= response.BaseExperience { // can catch
		color.Green("%s was caught!\n", pokemon)
		cfg.caughtPokemons[pokemon] = response 
	} else { 																		// cannot catch
		color.Red("%s excaped!\n", pokemon)
	}

	return nil
}