package shared

import (
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeApiClient   pokeapi.Client
	NextLocationUrl *string
	PrevLocationUrl *string
	CaughtPokemons  map[string]pokeapi.Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}
