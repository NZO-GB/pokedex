package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/NZO-GB/pokedex/internal/cache"
	"github.com/NZO-GB/pokedex/internal/pokeapi"
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

type configStruct struct {
	mapCall mapCall
	cache pokecache.Cache
	pokeClient *pokeapi.Client
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *configStruct) {
	scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			} else {
				fmt.Println("EOF, exiting REPL")
			}
			return
		}

			words := cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}

			command := words[0]

			cmd, exists := supportedCommands[command]
			if exists {
				fmt.Printf("Executing %s \n", command)
				if err := cmd.callback(cfg); err != nil{
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
}