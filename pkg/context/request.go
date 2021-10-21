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
	"io/ioutil"
	"net/http"

	"github.com/Codexiaoyi/linweb/interfaces"
)

var _ interfaces.IRequest = &Request{}

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
