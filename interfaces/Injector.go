package interfaces

type IInjector interface {
	AddSingleton(objects ...interface{})
	AddTransient(objects ...interface{})
	Inject(service interface{})
}
