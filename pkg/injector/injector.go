package injector

import (
	"linweb/interfaces"
	"reflect"
	"sync"
)

var _ interfaces.IInjector = &Injector{}

type mode int

const (
	singleton mode = iota
	transient
)

var (
	injectorOnce sync.Once
	injector     *Injector
)

type Injector struct {
	modeMap sync.Map
	sc      *singletonContainer
	tc      *transientContainer
}

func newInjector() *Injector {
	return &Injector{
		modeMap: sync.Map{},
		sc:      newSingletonContainer(),
		tc:      newTransientContainer(),
	}
}

// The injector can only be created once.
func Instance() interfaces.IInjector {
	injectorOnce.Do(func() {
		injector = newInjector()
	})
	return injector
}

func (ij *Injector) AddSingleton(objects ...interface{}) {
	if objects != nil && len(objects) != 0 {
		for _, object := range objects {
			ij.addObject(singleton, object)
		}
	}
}

func (ij *Injector) AddTransient(objects ...interface{}) {
	if objects != nil && len(objects) != 0 {
		for _, object := range objects {
			ij.addObject(transient, object)
		}
	}
}

func (ij *Injector) Inject(value reflect.Value) {
	if value.IsNil() || !value.IsValid() || value.IsZero() {
		return
	}
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if mode, ok := ij.modeMap.Load(field.Type); ok && value.Field(i).CanSet() {
			var result reflect.Value
			if mode == singleton {
				result = ij.sc.getObject(field.Type)
			} else {
				result = ij.tc.getObject(field.Type)
			}
			if !result.IsNil() || result.IsValid() || !result.IsZero() {
				value.Field(i).Set(result)
				ij.Inject(result)
			}
		}
	}
}

func (ij *Injector) addObject(mode mode, object interface{}) {
	// only add first registry object.
	if _, ok := ij.modeMap.Load(object); ok {
		return
	}
	ij.modeMap.Store(reflect.TypeOf(object), mode)
	// add a singleton object
	if mode == singleton {
		ij.sc.setObject(object)
	} else {
		ij.tc.setObject(object)
	}
}
