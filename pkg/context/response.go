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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Codexiaoyi/linweb/interfaces"
)

var _ interfaces.IResponse = &Response{}

type Response struct {
	response http.ResponseWriter
}

func NewResponse(res http.ResponseWriter) interfaces.IResponse {
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
