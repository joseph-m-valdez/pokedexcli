package repl

import (
	"fmt"
	"os"

	"github.com/joseph-m-valdez/pokedexcli/internal/config"
)

func commandExit(cfg *config.Config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
