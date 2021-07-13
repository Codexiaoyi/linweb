package context

import (
	"encoding/json"
	"fmt"
	"linweb/interfaces"
	"net/http"
)

type Response struct {
	response http.ResponseWriter
}

func newResponse(res http.ResponseWriter) interfaces.IResponse {
	return &Response{response: res}
}

func (res *Response) Status(code int) {
	res.response.WriteHeader(code)
}

func (res *Response) Header(key string, value string) {
	res.response.Header().Set(key, value)
}

func (res *Response) Data(code int, data []byte) {
	res.Status(code)
	res.response.Write(data)
}

func (res *Response) String(code int, format string, values ...interface{}) {
	res.Header("Content-Type", "text/plain")
	res.Data(code, []byte(fmt.Sprintf(format, values...)))
}

func (res *Response) JSON(code int, obj interface{}) {
	res.Header("Content-Type", "application/json")
	res.Status(code)
	encoder := json.NewEncoder(res.response)
	if err := encoder.Encode(obj); err != nil {
		http.Error(res.response, err.Error(), 500)
	}
}

func (res *Response) HTML(code int, html string) {
	res.Header("Content-Type", "text/plain")
	res.Data(code, []byte(html))
}
