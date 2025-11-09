package pokecache

import (
	"sync"
	"time"
) 

type Cache struct {
	locations map[string]CacheEntry
	mu sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time	
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache {
		locations: map[string]CacheEntry{},
		mu: sync.Mutex{},
		interval: interval,
	}
	go c.reapLoop()

	return &c
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<- ticker.C
		now := time.Now()
		c.mu.Lock()
		for key := range c.locations {
			if now.Sub(c.locations[key].createdAt) > c.interval {
				delete(c.locations, key)
			} 
		}
		c.mu.Unlock()
	}
}
