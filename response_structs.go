package main

import pokecache "github.com/NZO-GB/pokedex/internal/cache"

type mapCall struct {
	next string
	prev string
}

type configStruct struct {
	mapCall mapCall
	cache pokecache.Cache
}

type mapStruct struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configStruct) error
}

