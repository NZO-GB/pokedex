package main

import(
	"os"
	"fmt"
)

func commandExit(_ *configStruct, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
