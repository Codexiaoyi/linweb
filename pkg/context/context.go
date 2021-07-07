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

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		response: newResponse(w),
		request:  newRequest(req),
	}
}

func (c *Context) Request() interfaces.IRequest {
	return c.request
}

func (c *Context) Response() interfaces.IResponse {
	return c.response
}
