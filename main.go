package main

import (
	"time"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient 	pokeapi.Client
	nextLocationUrl *string 
	prevLocationUrl *string 
}

func main() {
	cfg := config{
		pokeapiClient: *pokeapi.NewClient(time.Hour),
	}

	startREPL(&cfg)
}