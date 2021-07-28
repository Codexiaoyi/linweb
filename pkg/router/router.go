package router

import (
	"encoding/json"
	"linweb/interfaces"
	"net/http"
	"reflect"
	"strings"
)

type Router struct {
	root     map[MethodType]*node
	handlers map[string]*Function
}

func New() interfaces.IRouter {
	return &Router{root: make(map[MethodType]*node), handlers: make(map[string]*Function)}
}

func (r *Router) AddControllers(controllers []interface{}) {
	parser := NewParser(controllers)
	for _, f := range parser.Funcs {
		r.addRoute(f.Method, f.Url, f)
	}
}

func (r *Router) addRoute(method MethodType, url string, handler *Function) {
	parts := parsePattern(url)
	key := getMethod(method) + "-" + url
	_, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
	}
	r.root[method].insert(url, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) getRoute(method MethodType, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.root[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.url)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *Router) Handle(c interfaces.IContext) {
	n, params := r.getRoute(getMethodType(c.Request().Method()), c.Request().Path())
	if n != nil {
		// set the params of url to the context.
		c.Request().SetParams(params)
		key := c.Request().Method() + "-" + n.url
		// map all route fuction to get the function info.
		handler := r.handlers[key]
		if handler != nil {
			// call controller's method.
			if !handler.Dto.IsValid() {
				handler.Recv.MethodByName(handler.Name).Call([]reflect.Value{reflect.ValueOf(c)})
			} else {
				// parse request body to map to the dto.
				parseJson(c.Request().Body(), handler.Dto)
				handler.Recv.MethodByName(handler.Name).Call([]reflect.Value{reflect.ValueOf(c), handler.Dto})
			}
			return
		}
	}
	// not has exists node in the trie, return 404.
	c.Response().String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Request().Path())
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// Parse request body to map to the dto.
func parseJson(js string, dto reflect.Value) {
	if js != "" {
		var fieldMap map[string]interface{}
		err := json.Unmarshal([]byte(js), &fieldMap)
		if err != nil {
			return
		}
		for i := 0; i < dto.Type().NumField(); i++ {
			field := fieldMap[dto.Type().Field(i).Name]
			if !dto.Field(i).CanSet() {
				// skip the field of can not set.
				continue
			}
			dto.Field(i).Set(reflect.ValueOf(field))
		}
	}
}
