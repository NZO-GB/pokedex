package main

import "fmt"

func commandHelp(config *configStruct) error {
	fmt.Println("Welcome to the Pokedex!\n\nUsage:\n")
	for command, information := range supportedCommands {
		fmt.Printf("%s:\t\t%s\n", command, information.description)
	}
	return nil
}
