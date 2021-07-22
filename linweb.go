package linweb

import (
	"linweb/interfaces"
	"linweb/pkg/context"
	"linweb/pkg/router"
	"net/http"
	"reflect"
)

type Linweb struct {
	router  interfaces.IRouter
	context interfaces.IContext
}

func New() *Linweb {
	return &Linweb{router: router.NewRouter(), context: context.NewContext()}
}

// You can import your IRouter and IContext implements.
func NewLinweb(router interfaces.IRouter, context interfaces.IContext) *Linweb {
	return &Linweb{router: router, context: context}
}

// Add all controllers, they must all be of pointer type
func (linweb *Linweb) AddControllers(obj ...interface{}) {
	linweb.router.AddControllers(obj)
}

// Run you project to listen the "addr", enjoy yourself!
func (linweb *Linweb) Run(addr string) error {
	return http.ListenAndServe(addr, linweb)
}

func (linweb *Linweb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//according the linweb's context implement to create a new implement type,
	//by call method "NewContext"
	newResults := reflect.New(reflect.ValueOf(linweb.context).Type()).Elem().MethodByName("New").Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(req)})
	//convert this object to interfaces.IContext
	context := newResults[0].Interface().(interfaces.IContext)
	linweb.router.Handle(context)
}
