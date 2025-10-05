package repl

import (
	"errors"

	"github.com/joseph-m-valdez/pokedexcli/internal/config"

	"fmt"
)

func commandExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	} 
	locationAreaName := args[0]

	locationArea, err := cfg.APIClient.ListLocationArea(locationAreaName)

	if err != nil {
		return err 
	}

	for _, pokemon := range locationArea.PokemonEncounters{
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
