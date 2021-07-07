package router

import (
	"linweb/interfaces"
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]interfaces.HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]interfaces.HandlerFunc)}
}

func (r *Router) addRoute(method string, part string, handler interfaces.HandlerFunc) {
	log.Printf("Route %4s - %s", method, part)
	key := method + "-" + part
	r.handlers[key] = handler
}

func (r *Router) Handle(c interfaces.IContext) {
	key := c.Request().Method() + "-" + c.Request().Path()
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.Response().String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Request().Path())
	}
}

func (r *Router) Get(part string, handler interfaces.HandlerFunc) {
	r.addRoute("GET", part, handler)
}

func (r *Router) Post(part string, handler interfaces.HandlerFunc) {
	r.addRoute("POST", part, handler)
}
