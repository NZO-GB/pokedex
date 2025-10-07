package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"math/rand"
)

func commandCatch(cfg *configStruct, pokemon string) error {

	if data, exists := cfg.pokedex[pokemon]; exists {
		catchPokemon(cfg, &data)
	}

	call := pokemonURL + pokemon

	req, err := http.NewRequest("GET", call, nil)
	if err != nil {
		return fmt.Errorf("NewReq found: %s", err)
	}

	res, err := cfg.pokeClient.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("ClientDo found: %s", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("ReadAll found: %s", err)
	}

	var jsStruct pokemonStruct

	if err := json.Unmarshal(data, &jsStruct); err != nil {
		fmt.Println(err)
		return fmt.Errorf("pokemon not found")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	catchPokemon(cfg, &jsStruct)

	return nil
}

func catchPokemon(cfg *configStruct, pokemonData *pokemonStruct) {
	pokemonExp := pokemonData.BaseExperience
	pokemonName := pokemonData.Name
	chance := rand.Intn(350)

	if chance >= pokemonExp {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = *pokemonData
	} else {
		fmt.Println("You missed!")
	}
}