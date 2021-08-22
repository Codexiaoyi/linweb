package interfaces

type Value interface {
	Len() int
}

// Memory cache. Must use type of implement Value as value of Cache.
// Because we need compute cache size by Value.Len().
type Cache interface {
	Get(key string) (value Value, ok bool)
	Add(key string, value Value)
	Delete(key string)
}
