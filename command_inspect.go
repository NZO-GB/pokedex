package main

import(
	"fmt"
)

func commandInspect(cfg *configStruct, pokemon string) error {

	data, exists := cfg.pokedex[pokemon]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	return pokemonStats(&data)
}

func pokemonStats(pokemon *pokemonStruct) error {
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	pokemonStats := make(map[string]int)
	for _, statStruct := range(pokemon.Stats) {
		pokemonStats[statStruct.Stat.Name] = statStruct.BaseStat

	}
	fmt.Println(" -hp: ", pokemonStats["hp"])
	fmt.Println(" -attack: ",pokemonStats["attack"])
	fmt.Println(" -defense: ", pokemonStats["defense"])
	fmt.Println(" -special-attack: ", pokemonStats["special-attack"])
	fmt.Println(" -special-defense: ", pokemonStats["special-defense"])
	fmt.Println(" -speed: ", pokemonStats["speed"])
	fmt.Println("Types:")
	for _, t := range(pokemon.Types) {
		fmt.Println(" - ", t.Type.Name)
	}
	
	return nil

}