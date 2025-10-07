package pokecache

import (
	"time"
	"sync"
	"fmt"
)
type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entryMap map[string]cacheEntry
	interval time.Duration
	mux *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache {
		entryMap:	make(map[string]cacheEntry),
		interval:	interval,
		mux:		&sync.Mutex{},
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.entryMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mux.Unlock()
	fmt.Printf("Added %s to cache\n", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, found := c.entryMap[key]
	c.mux.Unlock()
	fmt.Printf("Got %s from cache\n", key)
	return entry.val, found
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for key, value := range c.entryMap {
			if value.createdAt.Before(time.Now().Add(- c.interval)) {
				delete(c.entryMap, key)
			}
			
		}
		c.mux.Unlock()
	}
}