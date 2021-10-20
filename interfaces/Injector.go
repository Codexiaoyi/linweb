package interfaces

import "reflect"

// The DI container, regist object in the linweb init.
type IInjector interface {
	// Add singleton objects to the container.
	AddSingleton(objects ...interface{})
	// Add Transient objects to the container, create a new object in every reuqest.
	AddTransient(objects ...interface{})
	// Inject object in the request is coming.
	Inject(value reflect.Value)
}
