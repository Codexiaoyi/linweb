package interfaces

type HandlerFunc func(c IContext)

type IMiddleware interface {
	New(middlewareFuncs ...HandlerFunc) IMiddleware
	AddMiddlewares(middlewareFuncs ...HandlerFunc)
	Next(c IContext)
}
