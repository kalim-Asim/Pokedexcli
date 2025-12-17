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
	// exploreUrl       string 
}

func main() {
	cfg := config{
		pokeapiClient: *pokeapi.NewClient(time.Hour),
	}

	startREPL(&cfg)
}