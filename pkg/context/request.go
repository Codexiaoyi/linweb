package context

import "net/http"

type Request struct {
	req *http.Request
}

func newRequest(req *http.Request) *Request {
	return &Request{req: req}
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
