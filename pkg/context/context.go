package context

import (
	"linweb/interfaces"
	"net/http"
)

type Context struct {
	// origin objects
	response interfaces.IResponse
	request  interfaces.IRequest
	params   map[string]string
}

func NewContext(w http.ResponseWriter, req *http.Request) interfaces.IContext {
	return &Context{
		response: newResponse(w),
		request:  newRequest(req),
		params:   make(map[string]string),
	}
}

func (c *Context) Request() interfaces.IRequest {
	return c.request
}

func (c *Context) Response() interfaces.IResponse {
	return c.response
}

func (c *Context) SetParams(params map[string]string) {
	c.params = params
}
func (c *Context) Param(key string) string {
	return c.params[key]
}
