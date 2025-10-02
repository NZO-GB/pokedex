package main

import (
	"github.com/NZO-GB/pokedex/internal/cache"
	"time"
)

var cfg configStruct

func main() {
	cfg.cache = pokecache.NewCache(10 * time.Second)
	startRepl()
	
}
