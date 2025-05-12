package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{}
	newCache.entry = make(map[string]cacheEntry)
	newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.entry[key]
	if ok {
		return val.val, ok
	} else {
		return []byte{}, ok
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	if interval == 0 {
		fmt.Println("Cannot have a cache interval of 0.")
		return
	}
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			now := time.Now()
			cutoff := now.Add(-interval)

			c.mu.Lock()
			for key, entry := range c.entry {
				if entry.createdAt.Before(cutoff) {
					delete(c.entry, key)
					fmt.Printf("Clearing entry %v from cache at %v, cutoff time is %v\n", key, now, cutoff)
				}
			}
			c.mu.Unlock()
		}
	}()
}
