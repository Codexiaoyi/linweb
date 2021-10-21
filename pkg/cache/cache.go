//Copyright 2021 Codexiaoyi
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package cache

import (
	"sync"
	"time"

	"github.com/Codexiaoyi/linweb/interfaces"
)

var _ interfaces.ICache = &Cache{}

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

func newCache() *Cache {
	return &Cache{
		lru:     newLru(2, onKeyDelete),
		sweeper: newSweeper(5*time.Second, onKeyDelete),
	}
}

func onKeyDelete(key string) {
	if cache != nil {
		cache.Delete(key)
	}
}

// The cache can only be created once.
func Instance() interfaces.ICache {
	cacheOnce.Do(func() {
		cache = newCache()
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
