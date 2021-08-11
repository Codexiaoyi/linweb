package interfaces

import "net/http"

type IContext interface {
	// New create an instance
	New() IContext
	// Injecting dependencies and reset this Context. In order to use sync.Pool to reuse Context.
	Reset(http.ResponseWriter, *http.Request, IMiddleware)
	// Get http request info.
	Request() IRequest
	// Get http response info.
	Response() IResponse
	Middleware() IMiddleware
	Next()
}

// IRequest Get request data in *http.Request
type IRequest interface {
	// PostForm get post form
	PostForm(key string) string
	// Query get query data in url
	Query(key string) string
	Path() string
	Method() string
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
