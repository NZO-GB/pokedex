package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *configStruct) error {

	var call string

	if cfg.mapCall.next == "" {
		call = "https://pokeapi.co/api/v2/location-area/"
	} else {
		call = cfg.mapCall.next
	}
	if data, ok := cfg.cache.Get(call); ok {
		return commandMapPrinter(data, call, cfg, true)
	}

	return commandMapGetter(call, cfg, true)
}

func commandMapb(cfg *configStruct) error {
	if cfg.mapCall.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	if data, ok := cfg.cache.Get(cfg.mapCall.prev); ok {
		return commandMapPrinter(data, cfg.mapCall.prev, cfg, true)
	}

	return commandMapGetter(cfg.mapCall.prev, cfg, false)
}

func commandMapGetter(call string, cfg *configStruct, forward bool) error {

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

	commandMapPrinter(data, call, cfg, forward)

	return nil
}

func commandMapPrinter(data []byte, call string, cfg *configStruct, forward bool) error {

	cfg.cache.Add(call, data)

	var jsStruct mapStruct
	if err := json.Unmarshal(data, &jsStruct); err != nil {
		fmt.Printf("(Unmarshal) encountered error:\n, %s", err)
		return err
	}

	for _, location := range jsStruct.Results {
		fmt.Println(location.Name)
	}

	if forward {
		cfg.mapCall.prev = call
		cfg.mapCall.next = *jsStruct.Next
	} else {
		fmt.Printf("Setting prev as: %s\t Setting next as: %s\n", *jsStruct.Previous, call)
		cfg.mapCall.prev = *jsStruct.Previous
		cfg.mapCall.next = call
	}

	return nil

}
