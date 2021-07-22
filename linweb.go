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

func (linweb *Linweb) AddControllers(obj ...interface{}) {
	linweb.router.AddControllers(obj)
}

func (linweb *Linweb) Run(addr string) error {
	return http.ListenAndServe(addr, linweb)
}

func (linweb *Linweb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := context.NewContext(w, req)
	linweb.router.Handle(c)
}
