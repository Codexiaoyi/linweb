package interfaces

type IRouter interface {
	Handle(c IContext, i IInjector)
	AddControllers(controllers []interface{})
}
