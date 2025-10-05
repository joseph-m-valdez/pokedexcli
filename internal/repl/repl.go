// Package repl is the main function for the pokedex
package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joseph-m-valdez/pokedexcli/internal/config"
)

func StartRepl(cfg *config.Config) {
	cmds := getCommands()
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := cmds[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowerText)

	return strings.Fields(trimmed)
}

type cliCommand struct {
	name				string
	description	string
	callback		func(*config.Config, ...string ) error	
}

func getCommands() map[string]cliCommand {
	 return map[string]cliCommand{
		"exit": {
			name:					"exit",
			description:	"Exit the Pokedex",
			callback:			commandExit,
		},
		"help": {
			name:					"help",
			description:	"Displays a help message",
			callback:			commandHelp,
		},
		"map": {
			name:					"map",
			description: 	"Displays the next page of Poke locations",
			callback: 		commandMap,
		},
		"mapb": {
			name:					"mapb",
			description: 	"Displays the previous page of Poke locations",
			callback: 		commandMapb,
		},
		"explore": {
			name:					"explore {location area}",
			description: 	"Displays the pokemon in an area",
			callback: 		commandExplore,
		},
		"catch": {
			name:					"catch {pokemon name}",
			description: 	"Catch a pokemon and add it to your Pokedex",
			callback: 		commandCatch,
		},
		"inspect": {
			name:					"inspect {pokemon name}",
			description: 	"Inspect a pokemon in your pokedex.",
			callback: 		commandInspect,
		},
		"pokedex": {
			name:					"pokedex",
			description: 	"See all your caught pokemon",
			callback: 		commandPokedex,

		},
	}
}
