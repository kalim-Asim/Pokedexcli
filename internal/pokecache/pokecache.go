package pokecache

import (
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val				[]byte
}

type Cache struct {
	cache 	map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
    cache:  make(map[string]cacheEntry),
  }
	go c.reapLoop(interval)
  return c
}

func (c* Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: 			 val,
	}
}

func (c* Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key] 
	return cacheE.val, ok
}

func (c* Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-1 * interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}

func (c* Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}