package router

import (
	"linweb/interfaces"
	"net/http"
	"strings"
)

type Router struct {
	root     map[string]*node
	handlers map[string]interfaces.HandlerFunc
}

func NewRouter() *Router {
	return &Router{root: make(map[string]*node), handlers: make(map[string]interfaces.HandlerFunc)}
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

func (r *Router) addRoute(method string, url string, handler interfaces.HandlerFunc) {
	parts := parsePattern(url)

	key := method + "-" + url
	_, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
	}
	r.root[method].insert(url, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) Handle(c interfaces.IContext) {
	n, params := r.getRoute(c.Request().Method(), c.Request().Path())
	if n != nil {
		c.SetParams(params)
		key := c.Request().Method() + "-" + c.Request().Path()
		r.handlers[key](c)
	} else {
		c.Response().String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Request().Path())
	}
}

func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
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

func (r *Router) Get(url string, handler interfaces.HandlerFunc) {
	r.addRoute("GET", url, handler)
}

func (r *Router) Post(url string, handler interfaces.HandlerFunc) {
	r.addRoute("POST", url, handler)
}
