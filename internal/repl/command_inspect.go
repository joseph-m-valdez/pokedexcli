package repl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/joseph-m-valdez/pokedexcli/internal/config"
)


func commandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	
	pokemon, exists := cfg.CaughtPokemon[pokemonName]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Name: %s \n", pokemon.Name))
	builder.WriteString(fmt.Sprintf("Height: %d \n", pokemon.Height))
	builder.WriteString(fmt.Sprintf("Weight: %d \n", pokemon.Weight))
	builder.WriteString("Stats:\n")
	for _, s := range pokemon.Stats {
		builder.WriteString(fmt.Sprintf("  -%s: %d\n", s.Stat.Name, s.BaseStat))
	}
	builder.WriteString("Types:\n")
	for _, t := range pokemon.Types {
		builder.WriteString(fmt.Sprintf("  - %s\n", t.Type.Name))
	}


	pokemonInfo := builder.String()

	fmt.Println(pokemonInfo)


	return nil
}
