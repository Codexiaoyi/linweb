package interfaces

type HandlerFunc func(c IContext)

type IMiddleware interface {
	// New create an instance
	New(middlewareFunc ...HandlerFunc) IMiddleware
	// Add middlewares to the request context.
	AddMiddlewares(middlewareFunc ...HandlerFunc)
	// Call the next middleware.
	Next(c IContext)
}
