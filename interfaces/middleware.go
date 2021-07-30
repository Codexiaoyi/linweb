package interfaces

type HandlerFunc func(c IContext)

type IMiddleware interface {
	New(middlewareFunc ...HandlerFunc) IMiddleware
	AddMiddlewares(middlewareFunc ...HandlerFunc)
	Next(c IContext)
}
