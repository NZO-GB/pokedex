package main

import(
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func commandMap(config *configStruct) error {
	if config.mapCall.next == "" {
		return commandMapHelper("https://pokeapi.co/api/v2/location-area/", true)
	}
	return commandMapHelper(config.mapCall.next, true)
}

func commandMapb(config *configStruct) error {
	if config.mapCall.prev == "" || config.mapCall.prev == "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20" {
	fmt.Println("you're on the first page")
		return nil
	}
	return commandMapHelper(config.mapCall.prev, false)
}

func commandMapHelper(call string, forward bool) error {
	res, err := http.Get(call)
	if err != nil {
		fmt.Printf("(Get) encountered error:\n, %s", err)
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("(ReadAll) encountered error:\n, %s", err)
		return err
	}

	var jsStruct mapStruct
	if err := json.Unmarshal(data, &jsStruct); err != nil {
		fmt.Printf("(Unmarshal) encountered error:\n, %s", err)
		return err
	}

	for _, location := range jsStruct.Results {
		fmt.Println(location.Name)
	}

	if forward {
		config.mapCall.prev = call
		config.mapCall.next = *jsStruct.Next
	} else {
		config.mapCall.prev = *jsStruct.Previous
		config.mapCall.next = call
	}

	return nil
}