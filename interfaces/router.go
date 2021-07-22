package interfaces

type HandlerFunc func(IContext)

type IRouter interface {
	Handle(c IContext)
	AddControllers(controllers []interface{})
}
