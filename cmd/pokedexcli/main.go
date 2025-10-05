// Package pokedexcli main entrypoint
package pokedexcli

import (
	"time"

	"github.com/joseph-m-valdez/pokedexcli/internal/api"
	"github.com/joseph-m-valdez/pokedexcli/internal/config"
	"github.com/joseph-m-valdez/pokedexcli/internal/repl"
)

func Main() {
	cfg := config.Config {
		CaughtPokemon: map[string]api.Pokemon{},
		APIClient: api.NewClient(time.Hour),
	}

	repl.StartRepl(&cfg)
}
