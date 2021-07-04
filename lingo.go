package lingo

import (
	"lingo/interfaces"
	"lingo/pkg/context"
	"lingo/pkg/router"
	"net/http"
)

type Lingo struct {
	router interfaces.IRouter
}

func New() *Lingo {
	return &Lingo{router: router.NewRouter()}
}

func (lingo *Lingo) GET(part string, handler interfaces.HandlerFunc) {
	lingo.router.Get(part, handler)
}

func (lingo *Lingo) POST(part string, handler interfaces.HandlerFunc) {
	lingo.router.Post(part, handler)
}

func (lingo *Lingo) Run(addr string) error {
	return http.ListenAndServe(addr, lingo)
}

func (lingo *Lingo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := context.NewContext(w, req)
	lingo.router.Handle(c)
}
