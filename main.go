package main

import (
	"github.com/NZO-GB/pokedex/internal/pokeapi"
	"github.com/NZO-GB/pokedex/internal/cache"
	"time"
)

func main() {
	
	var cfg configStruct
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg.pokeClient = &pokeClient
	cfg.cache = pokecache.NewCache(10 * time.Second)
	cfg.mapCall.next = &locationURL
	cfg.pokedex = make(map[string]pokemonStruct)
	startRepl(&cfg)
}
