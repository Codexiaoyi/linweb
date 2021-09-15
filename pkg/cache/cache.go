package cache

import (
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
	mux      sync.RWMutex
}

// The cache can only be created once.
func Instance() *Cache {
	cacheOnce.Do(func() {
		onKeyDelete := func(key string) {
			if cache != nil {
				cache.Delete(key)
			}
		}
		cache = &Cache{
			lru:     newLru(2, onKeyDelete),
			sweeper: newSweeper(5*time.Second, onKeyDelete),
		}
	})
	return cache
}

// Set the cache can use max bytes.
func (c *Cache) SetMaxBytes(maxBytes int64) {
	c.maxBytes = maxBytes
}

// Get value by key.If the key is expired, return "nil,false".
func (c *Cache) Get(key string) (value interface{}, ok bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	// expired
	if c.sweeper.isExpired(key) {
		c.lru.delete(key)
		c.sweeper.delete(key, false)
		return nil, false
	}
	return c.lru.get(key)
}

func add(key string, value interface{}) {
	cache.lru.add(key, value)
	for cache.maxBytes != 0 && cache.maxBytes < cache.lru.currentBytes {
		cache.lru.removeOldest()
	}
}

// Add cache.
func (c *Cache) Add(key string, value interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	add(key, value)
}

// Add cache with expire time.
func (c *Cache) AddWithExpire(key string, value interface{}, expireDuration time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.sweeper.addExpireKey(key, expireDuration)
	add(key, value)
}

// Delete cache.
func (c *Cache) Delete(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.lru.delete(key)
	c.sweeper.delete(key, false)
}
