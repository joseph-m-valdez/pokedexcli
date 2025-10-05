package repl

import (
	"fmt"

	"github.com/joseph-m-valdez/pokedexcli/internal/config"
)

func commandPokedex(cfg *config.Config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
