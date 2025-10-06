package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"

)

func commandExplore(cfg *configStruct, area string) error {
	fmt.Printf("Exploring %s...", area)
	areaURL := locationURL + area

	if data, ok := cfg.cache.Get(areaURL); ok {
		return areaPrinter(cfg, data, areaURL)
	} else {
		return areaGetter(cfg, areaURL)
	}
}

func areaGetter(cfg *configStruct, areaURL string) error {

	req, err := http.NewRequest("GET", areaURL, nil)
	if err != nil {
		return fmt.Errorf("area NewReq found: %s", err)
	} 

	res, err := cfg.pokeClient.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("area ClientDo found: %s", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("area ReadAll found: %s", err)
	}

	return areaPrinter(cfg, data, areaURL)
}

func areaPrinter(cfg *configStruct, data []byte, call string) error {
	
	cfg.cache.Add(call, data)

	var jsStruct areaStruct
	if err := json.Unmarshal(data, &jsStruct); err != nil {
		return fmt.Errorf("area Unmarshal found: %s", err)
	}

	fmt.Println("Found Pokémon:")
	for _, encounter := range(jsStruct.PokemonEncounters) {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}	