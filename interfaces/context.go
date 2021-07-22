package interfaces

import "net/http"

type IContext interface {
	//create an instance
	New(http.ResponseWriter, *http.Request) IContext
	//set params in url
	SetParams(map[string]string)
	Request() IRequest
	Response() IResponse
	//get url param
	Param(string) string
}

// Get request data in *http.Request
type IRequest interface {
	// get post form
	PostForm(key string) string
	// get query data
	Query(key string) string
	Path() string
	Method() string
	Body() string
}

// Set response data to http.ResponseWriter
type IResponse interface {
	// set response status
	Status(code int)
	// set response header
	Header(key string, value string)
	// set response body data
	Data(code int, data []byte)
	// set response body data with string
	String(code int, format string, values ...interface{})
	// set response body data with json
	JSON(code int, obj interface{})
	// set response body data with html
	HTML(code int, html string)
}