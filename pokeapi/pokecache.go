package pokeapi

import "time"

type cacheEntry struct {
    createdAt time.Time
    val       []byte
}

type Cache struct {
    cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
    c := Cache{
        cache: make(map[string]cacheEntry),
    }
    go c.reapLoop(interval)
    return c
}

func (c *Cache) Add(key string, value []byte) {
    c.cache[key] = cacheEntry{
        val:       value,
        createdAt: time.Now(),
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    if currentCacheEntry, ok := c.cache[key]; ok {
        return currentCacheEntry.val, true
    }

    return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)

    defer ticker.Stop()

    for range ticker.C {
        now := time.Now()
        for key, entry := range c.cache {
            if now.Sub(entry.createdAt) > interval {
                delete(c.cache, key)
            }
        }
    }
}
