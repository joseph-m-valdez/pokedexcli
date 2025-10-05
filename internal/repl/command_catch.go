package repl

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/joseph-m-valdez/pokedexcli/internal/config"
)


func commandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.APIClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)	
	if pokemon.BaseExperience / 2 >= rand.IntN(pokemon.BaseExperience) {
		fmt.Println("You caught it!")
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} 

	return nil
}
