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

package interfaces

import "net/http"

// IContext Save http request data.
type IContext interface {
	// New create an instance
	New() IContext
	// Injecting dependencies and reset this Context. In order to use sync.Pool to reuse Context.
	Reset(http.ResponseWriter, *http.Request, IMiddleware)
	// Get http request info.
	Request() IRequest
	// Get http response info.
	Response() IResponse
	// Get middleware in context.
	Middleware() IMiddleware
	// Next middleware.
	Next()
}

// IRequest Get request data in *http.Request
type IRequest interface {
	// PostForm get post form
	PostForm(key string) string
	// Query get query data in url
	Query(key string) string
	// Get url path.
	Path() string
	// Get http method.
	Method() string
	// Get request body.
	Body() string
	// SetParams set params in url
	SetParams(map[string]string)
	// Param get param in url
	Param(string) string
}

// IResponse Set response data to http.ResponseWriter
type IResponse interface {
	// Status set response status
	Status(code int)
	// Header set response header
	Header(key string, value string)
	// Data set response body data
	Data(code int, data []byte)
	// String set response body data with string
	String(code int, format string, values ...interface{})
	// JSON set response body data with json
	JSON(code int, obj interface{})
	// HTML set response body data with html
	HTML(code int, html string)
}
