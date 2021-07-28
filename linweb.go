package linweb

import (
	"fmt"
	"linweb/interfaces"
	"linweb/pkg/context"
	"linweb/pkg/router"
	"log"
	"net/http"
	"reflect"
)

type Linweb struct {
	router  interfaces.IRouter
	context interfaces.IContext
}

func NewLinweb() *Linweb {
	return &Linweb{}
}

// Add customize plugins, they must all be of pointer type.
// It is not allowed to pass in non-plugin implementations.
// Without customize plugins will use the default plugins.
func (linweb *Linweb) AddCustomizePlugins(plugins ...interface{}) {
	for _, p := range plugins {
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IRouter)(nil)).Elem()) {
			linweb.router = p.(interfaces.IRouter)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IContext)(nil)).Elem()) {
			linweb.context = p.(interfaces.IContext)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IModel)(nil)).Elem()) {
			plugins_model = p.(interfaces.IModel)
			continue
		}
		log.Fatal(fmt.Sprintf("'%s' is not a plugin, please check if it implements a plugin", reflect.TypeOf(p).Elem().Name()))
	}
}

// Add all controllers, they must all be of pointer type
func (linweb *Linweb) AddControllers(obj ...interface{}) {
	if linweb.router == nil {
		linweb.router = router.New()
	}
	linweb.router.AddControllers(obj)
}

// Run you project to listen the "addr", enjoy yourself!
func (linweb *Linweb) Run(addr string) error {
	return http.ListenAndServe(addr, linweb)
}

func (linweb *Linweb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if linweb.context == nil {
		linweb.context = &context.Context{}
	}
	//create a new context for current request
	context := linweb.context.New(w, req)
	linweb.router.Handle(context)
}
