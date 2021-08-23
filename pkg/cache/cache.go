package cache

import (
	"linweb/interfaces"
	"sync"
	"time"
)

var (
	cacheOnce sync.Once
	cache     *Cache
)

type Cache struct {
	// max bytes can use,
	// if maxBytes == 0, it will take up unlimited memory space.
	maxBytes int64
	lru      *lru
	sweeper  *sweeper
	mux      sync.Mutex
}

// The cache can only be created once.
func New(maxBytes int64) *Cache {
	cacheOnce.Do(func() {
		onExpireDelete := func(key string) {
			if cache != nil {
				cache.Delete(key)
			}
		}
		cache = &Cache{
			maxBytes: maxBytes,
			lru:      newLru(2),
			sweeper:  newSweeper(5, onExpireDelete),
		}
	})
	return cache
}

func (c *Cache) Get(key string) (value interfaces.Value, ok bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.lru.get(key)
}

func add(key string, value interfaces.Value) {
	cache.lru.add(key, value)
	for cache.maxBytes != 0 && cache.maxBytes < cache.lru.currentBytes {
		cache.lru.removeOldest()
	}
}

func (c *Cache) Add(key string, value interfaces.Value) {
	c.mux.Lock()
	defer c.mux.Unlock()
	add(key, value)
}

func (c *Cache) AddWithExpire(key string, value interfaces.Value, expireDuration time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.sweeper.addExpireKey(key, expireDuration)
	add(key, value)
}

func (c *Cache) Delete(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.lru.delete(key)
}
