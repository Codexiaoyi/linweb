package context

import (
	"linweb/interfaces"
	"net/http"
)

type Context struct {
	// origin objects
	response   interfaces.IResponse
	request    interfaces.IRequest
	middleware interfaces.IMiddleware
}

func (c *Context) New(w http.ResponseWriter, req *http.Request, m interfaces.IMiddleware) interfaces.IContext {
	return &Context{
		response:   NewResponse(w),
		request:    NewRequest(req),
		middleware: m,
	}
}

func (c *Context) Request() interfaces.IRequest {
	return c.request
}

func (c *Context) Response() interfaces.IResponse {
	return c.response
}

func (c *Context) Middleware() interfaces.IMiddleware {
	return c.middleware
}

func (c *Context) Next() {
	c.middleware.Next(c)
}
