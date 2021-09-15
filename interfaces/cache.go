package interfaces

import "time"

// Memory cache.
type ICache interface {
	Get(key string) (value interface{}, ok bool)
	Add(key string, value interface{})
	AddWithExpire(key string, value interface{}, expireDuration time.Duration)
	Delete(key string)
}
