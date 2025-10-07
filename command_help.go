package main

import "fmt"

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"exit": 	{name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help": 	{name: "help", description: "Displays a help message", callback: commandHelp},
		"map":  	{name: "map", description: "Presents the list of the next 20 location-areas", callback: commandMap},
		"mapb": 	{name: "mapb", description: "Presents the list of the previous 20 location-areas", callback: commandMapb},
		"explore": 	{name: "explore", description: "Presents the pokemon present in an given area", callback: commandExplore},
		"catch":	{name: "catch", description: "Attempts to catch a pokemon to add it to pokedex", callback: commandCatch},
		"inspect":	{name: "inspect", description: "Inspects pokemon in your pokedex", callback: commandInspect},
	}
}


func commandHelp(_ *configStruct, _ string) error {
	fmt.Println("Welcome to the Pokedex!\n\nUsage:")
	for command, information := range supportedCommands {
		fmt.Printf("%s:\t\t%s\n", command, information.description)
	}
	return nil
}
