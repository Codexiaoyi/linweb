//Copyright 2021 Codexiaoyi
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package context

import (
	"github.com/Codexiaoyi/linweb/interfaces"
	"net/http"
)

var _ interfaces.IContext = &Context{}

type Context struct {
	// origin objects
	response   interfaces.IResponse
	request    interfaces.IRequest
	middleware interfaces.IMiddleware
}

func (c *Context) New() interfaces.IContext {
	return &Context{}
}

func (c *Context) Reset(w http.ResponseWriter, req *http.Request, m interfaces.IMiddleware) {
	c.request = NewRequest(req)
	c.response = NewResponse(w)
	c.middleware = m
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
