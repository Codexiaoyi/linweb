package interfaces

type IRouter interface {
	// Handle the request.
	Handle(c IContext, i IInjector)
	// Add controllers to the router.
	AddControllers(controllers []interface{})
}
