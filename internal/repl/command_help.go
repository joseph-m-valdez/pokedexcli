package repl

import (
	"fmt"

	"github.com/joseph-m-valdez/pokedexcli/internal/config"
)

func commandHelp(cfg *config.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)	
	}
	fmt.Println()
	return nil
}
