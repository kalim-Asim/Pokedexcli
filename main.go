package main

import (
	"time"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

const baseURL = `https://pokeapi.co/api/v2`

type config struct {
	pokeapiClient 	pokeapi.Client
	nextLocationUrl *string 
	prevLocationUrl *string 
	caughtPokemons map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:  *pokeapi.NewClient(time.Hour),
		caughtPokemons:  make(map[string]pokeapi.Pokemon),
	}

	startREPL(&cfg)
}