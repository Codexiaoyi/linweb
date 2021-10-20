package middleware

import (
	"linweb/interfaces"
)

var _ interfaces.IMiddleware = &Middleware{}

type Middleware struct {
	funcs []interfaces.HandlerFunc
	index int
}

func (m *Middleware) New(middlewareFunc ...interfaces.HandlerFunc) interfaces.IMiddleware {
	return &Middleware{
		funcs: middlewareFunc,
		index: -1,
	}
}

func (m *Middleware) Next(context interfaces.IContext) {
	m.index++
	length := len(m.funcs)
	for ; m.index < length; m.index++ {
		m.funcs[m.index](context)
	}
}

func (m *Middleware) AddMiddlewares(middlewareFunc ...interfaces.HandlerFunc) {
	m.funcs = append(m.funcs, middlewareFunc...)
}
