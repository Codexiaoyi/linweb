package linweb

import (
	"linweb/interfaces"
	"linweb/pkg/context"
	"linweb/pkg/router"
	"net/http"
)

type Linweb struct {
	router interfaces.IRouter
}

func New() *Linweb {
	return &Linweb{router: router.NewRouter()}
}

func (linweb *Linweb) GET(part string, handler interfaces.HandlerFunc) {
	linweb.router.Get(part, handler)
}

func (linweb *Linweb) POST(part string, handler interfaces.HandlerFunc) {
	linweb.router.Post(part, handler)
}

func (linweb *Linweb) Run(addr string) error {
	return http.ListenAndServe(addr, linweb)
}

func (linweb *Linweb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := context.NewContext(w, req)
	linweb.router.Handle(c)
}
