package cache

import (
	"linweb/interfaces"
	"sync"
)

type Cache struct {
	// max bytes can use,
	// if maxBytes == 0, it will take up unlimited memory space.
	maxBytes int64
	lru      *lru
	mux      sync.Mutex
}

func New(maxBytes int64) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		lru:      newLru(2),
	}
}

func (c *Cache) Get(key string) (value interfaces.Value, ok bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.lru.get(key)
}

func (c *Cache) Add(key string, value interfaces.Value) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.lru.add(key, value)
	for c.maxBytes != 0 && c.maxBytes < c.lru.currentBytes {
		c.lru.removeOldest()
	}
}

func (c *Cache) Delete(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.lru.delete(key)
}
