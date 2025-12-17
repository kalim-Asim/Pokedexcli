package main

import (
	"time"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
	"github.com/kalim-Asim/pokedexcli/shared"
)

func main() {
	cfg := shared.Config{
		PokeApiClient:  *pokeapi.NewClient(time.Hour),
		CaughtPokemons:  make(map[string]pokeapi.Pokemon),
	}

	startREPL(&cfg)
}