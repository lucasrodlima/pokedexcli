package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	tick := time.Tick(interval)

	for {
		<-tick
		c.mu.Lock()

		now := time.Now()

		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}

		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	newC := &Cache{
		entries: map[string]cacheEntry{},
	}

	go newC.reapLoop(interval)

	return newC
}
