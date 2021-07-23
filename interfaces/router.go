package interfaces

type IRouter interface {
	Handle(c IContext)
	AddControllers(controllers []interface{})
}
