package interfaces

import "time"

// Memory cache.
type ICache interface {
	// Set the cache can use max bytes.
	SetMaxBytes(maxBytes int64)
	// Get value by key.If the key is expired, return "nil,false".
	Get(key string) (value interface{}, ok bool)
	// Add cache key.
	Add(key string, value interface{})
	// Add cache key with expire time.
	AddWithExpire(key string, value interface{}, expireDuration time.Duration)
	// Delete cache key.
	Delete(key string)
}
