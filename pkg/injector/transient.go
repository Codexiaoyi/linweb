package injector

import (
	"reflect"
	"sync"

	"github.com/huandu/go-clone"
)

// The transient object container.
type transientContainer struct {
	transient sync.Map
}

func newTransientContainer() *transientContainer {
	return &transientContainer{
		transient: sync.Map{},
	}
}

func (s *transientContainer) setObject(object interface{}) {
	t := reflect.TypeOf(object)
	if t.Kind() != reflect.Ptr {
		panic(t.Name() + " not ptr type.")
	}
	s.transient.Store(t, reflect.ValueOf(object))
}

func (s *transientContainer) getObject(t reflect.Type) reflect.Value {
	if value, ok := s.transient.Load(t); ok {
		return s.getNewObj(value.(reflect.Value))
	}
	res := reflect.Value{}
	//maybe object type is interface
	s.transient.Range(func(key, value interface{}) bool {
		if key.(reflect.Type).Kind() == reflect.Interface && key.(reflect.Type).Implements(t) {
			res = s.getNewObj(value.(reflect.Value))
			return false
		}
		return true
	})
	return res
}

func (s *transientContainer) getNewObj(oldObj reflect.Value) reflect.Value {
	obj := clone.Clone(oldObj.Interface())
	return reflect.ValueOf(obj)
}
