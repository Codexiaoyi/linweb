package injector

import (
	"reflect"
	"sync"
)

// The singleton object container.
type singletonContainer struct {
	singleton sync.Map
}

func newSingletonContainer() *singletonContainer {
	return &singletonContainer{
		singleton: sync.Map{},
	}
}

func (s *singletonContainer) setObject(object interface{}) {
	t := reflect.TypeOf(object)
	if t.Kind() != reflect.Ptr {
		panic(t.Name() + " not ptr type.")
	}
	s.singleton.Store(t, reflect.ValueOf(object))
}

func (s *singletonContainer) getObject(t reflect.Type) reflect.Value {
	if value, ok := s.singleton.Load(t); ok {
		return value.(reflect.Value)
	}
	res := reflect.Value{}
	//maybe object type is interface
	s.singleton.Range(func(key, value interface{}) bool {
		if key.(reflect.Type).Kind() == reflect.Interface && key.(reflect.Type).Implements(t) {
			res = value.(reflect.Value)
			return false
		}
		return true
	})
	return res
}
