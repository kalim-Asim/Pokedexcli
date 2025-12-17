package pokeapi

import (
	"net/http"
	"time"

	"github.com/kalim-Asim/pokedexcli/internal/pokecache"
	"github.com/kalim-Asim/pokedexcli/shared"
)

var BaseURL = `https://pokeapi.co/api/v2`

type Client struct {
	Cache      pokecache.Cache
	HttpClient http.Client
}

type Config struct {
	PokeApiClient   Client
	NextLocationUrl *string
	PrevLocationUrl *string
	CaughtPokemons  map[string]shared.Pokemon
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		Cache: 			pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}