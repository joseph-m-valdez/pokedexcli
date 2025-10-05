// Package config for pokeapi configuration
package config

import "github.com/joseph-m-valdez/pokedexcli/internal/api"

type Config struct {
  APIClient       api.Client
  NextLocationURL *string
  PrevLocationURL *string
	CaughtPokemon		map[string]api.Pokemon
}

