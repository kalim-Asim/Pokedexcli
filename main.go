package main

import (
	"time"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
	"github.com/kalim-Asim/pokedexcli/shared"
)

func main() {
	cfg := pokeapi.Config{
		PokeApiClient:  *pokeapi.NewClient(time.Hour),
		CaughtPokemons:  make(map[string]shared.Pokemon),
	}

	startREPL(&cfg)
}