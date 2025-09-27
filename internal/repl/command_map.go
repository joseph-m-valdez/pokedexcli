package repl

import (
	"github.com/joseph-m-valdez/pokedexcli/internal/config"

	"fmt"
	"errors"
)

func commandMap(cfg *config.Config) error {
	resp, err := cfg.APIClient.ListLocationAreas(cfg.NextLocationURL)

	if err != nil {
		return fmt.Errorf("ruh-Roh %w", err)
	}

	for _, area := range resp.Results{
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.NextLocationURL = resp.Next
	cfg.PrevLocationURL = resp.Previous

	return nil
}

func commandMapb(cfg *config.Config) error {
	if cfg.PrevLocationURL == nil {
		return errors.New("your'e on the first page")
	}

	resp, err := cfg.APIClient.ListLocationAreas(cfg.PrevLocationURL)

	if err != nil {
		return fmt.Errorf("ruh-Roh %w", err)
	}

	for _, area := range resp.Results{
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.NextLocationURL = resp.Next
	cfg.PrevLocationURL = resp.Previous

	return nil
}
