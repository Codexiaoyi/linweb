package context

import (
	"io/ioutil"
	"linweb/interfaces"
	"net/http"
)

type Request struct {
	req    *http.Request
	params map[string]string
}

func NewRequest(req *http.Request) interfaces.IRequest {
	return &Request{
		req:    req,
		params: make(map[string]string),
	}
}

func (req *Request) PostForm(key string) string {
	return req.req.FormValue(key)
}

func (req *Request) Query(key string) string {
	return req.req.URL.Query().Get(key)
}

func (req *Request) Path() string {
	return req.req.URL.Path
}

func (req *Request) Method() string {
	return req.req.Method
}

func (req *Request) Body() string {
	body, err := ioutil.ReadAll(req.req.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func (req *Request) SetParams(params map[string]string) {
	req.params = params
}

func (req *Request) Param(key string) string {
	return req.params[key]
}
