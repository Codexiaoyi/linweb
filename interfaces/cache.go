package interfaces

import "time"

type Value interface {
	Len() int
}

// Memory cache. Must use type of implement Value as value of Cache.
// Because we need compute cache size by Value.Len().
type ICache interface {
	Get(key string) (value Value, ok bool)
	Add(key string, value Value)
	AddWithExpire(key string, value Value, expireDuration time.Duration)
	Delete(key string)
}
