package context

import (
	"linweb/interfaces"
	"net/http"
)

type Context struct {
	// origin objects
	response interfaces.IResponse
	request  interfaces.IRequest
}

func (c *Context) New(w http.ResponseWriter, req *http.Request) interfaces.IContext {
	return &Context{
		response: NewResponse(w),
		request:  NewRequest(req),
	}
}

func (c *Context) Request() interfaces.IRequest {
	return c.request
}

func (c *Context) Response() interfaces.IResponse {
	return c.response
}
