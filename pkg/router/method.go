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

package router

import "strings"

type MethodType int

const (
	Unknown MethodType = iota
	GET
	POST
	PUT
	DELETE
	PATCH
	OPTIONS
	HEAD
)

// Get method type by method string.
func getMethodType(method string) MethodType {
	switch strings.ToUpper(method) {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	case "PATCH":
		return PATCH
	case "OPTIONS":
		return OPTIONS
	case "HEAD":
		return HEAD
	default:
		return Unknown
	}
}

// Get method string by method type.
func getMethod(t MethodType) string {
	switch t {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case PATCH:
		return "PATCH"
	case OPTIONS:
		return "OPTIONS"
	case HEAD:
		return "HEAD"
	default:
		return ""
	}
}
