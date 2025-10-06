package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func callGetter(cfg *configStruct, forward bool) (string, bool) {
	var call string
	switch forward {
	case true:
		if cfg.mapCall.next == nil {
			fmt.Println("you're on the last page")
			return "", false
	 	}else {
			call = *cfg.mapCall.next
		}
	case false:
		if cfg.mapCall.prev == nil {
			fmt.Println("you're on the first page")
			return "", false
		}else {
			call = *cfg.mapCall.prev
		}
	}
	return call, true
}

func commandMap(cfg *configStruct, _ string) error {
	call, proceed := callGetter(cfg, true)
	if proceed {
		fmt.Printf("Got call %s \n", call)
	} else {
		return nil
	}
	if data, ok := cfg.cache.Get(call); ok {
		fmt.Println("Got cache")
		return mapPrinter(data, call, cfg)
	}

	return mapGetter(call, cfg)
}

func commandMapb(cfg *configStruct, _ string) error {
	call, proceed := callGetter(cfg, false)
	if proceed {
		fmt.Printf("Got call %s \n", call)
	} else {
		return nil
	}
	if data, ok := cfg.cache.Get(call); ok {
		fmt.Println("Got cache")
		return mapPrinter(data, call, cfg)
	} 
	return mapGetter(call, cfg)
}

func mapGetter(call string, cfg *configStruct) error {

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
		return fmt.Errorf("(ReadAll) encountered error:\n, %s", err)
	}

	return mapPrinter(data, call, cfg)
}

func mapPrinter(data []byte, call string, cfg *configStruct) error {

	cfg.cache.Add(call, data)

	var jsStruct mapStruct
	if err := json.Unmarshal(data, &jsStruct); err != nil {
		return fmt.Errorf("(Unmarshal) encountered error:\n, %s", err)
	}

	for _, location := range jsStruct.Results {
		fmt.Println(location.Name)
	}
		cfg.mapCall.prev = jsStruct.Previous
		cfg.mapCall.next = jsStruct.Next

	return nil	

}
