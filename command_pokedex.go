package main

import "fmt"

func commandPokedex(cfg *configStruct, arg string) error {

	fmt.Println("Your Pokedex:")

	for pokemon := range(cfg.pokedex) {
		fmt.Println(" - ", pokemon)
	}

	return nil
}