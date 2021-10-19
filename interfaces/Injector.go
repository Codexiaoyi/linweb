package interfaces

import "reflect"

type IInjector interface {
	AddSingleton(objects ...interface{})
	AddTransient(objects ...interface{})
	Inject(value reflect.Value)
}
