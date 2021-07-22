package router

import "strings"

type MethodType int

const (
	Unknown MethodType = iota
	GET
	POST
	PUT
	DELETE
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
	default:
		return ""
	}
}
