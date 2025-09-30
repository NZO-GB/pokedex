package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help": {name: "help", description: "Displays a help message", callback: commandHelp},
		"map":  {name: "map", description: "Presents the list of the next 20 location-areas", callback: commandMap},
		"mapb": {name: "mapb", description: "Presents the list of the previous 20 location-areas", callback: commandMapb},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
			scanner.Scan()

			words := cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}

			command := words[0]

			cmd, ok := supportedCommands[command]
			if !ok {
				fmt.Println("Unknown command")
			}

			cmd.callback(&config)
		}
}