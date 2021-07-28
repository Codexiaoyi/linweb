package linweb

import (
	"fmt"
	"linweb/interfaces"
	"linweb/pkg/context"
	"linweb/pkg/middleware"
	"linweb/pkg/model"
	"linweb/pkg/router"
	"log"
	"net/http"
	"reflect"
)

type Linweb struct {
	router       interfaces.IRouter
	mark_context interfaces.IContext
	// this middleware means an implement, every request need create a new middleware from New() by mark_middleware.
	mark_middleware interfaces.IMiddleware
	middlewareFuncs []interfaces.HandlerFunc
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
			linweb.mark_context = p.(interfaces.IContext)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IModel)(nil)).Elem()) {
			plugins_model = p.(interfaces.IModel)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IMiddleware)(nil)).Elem()) {
			linweb.mark_middleware = p.(interfaces.IMiddleware)
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

// Add global middlewares
func (linweb *Linweb) AddMiddlewares(middlewareFuncs ...interfaces.HandlerFunc) {
	linweb.middlewareFuncs = middlewareFuncs
}

// Run you project to listen the "addr", enjoy yourself!
func (linweb *Linweb) Run(addr string) error {
	return http.ListenAndServe(addr, linweb)
}

func (linweb *Linweb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if linweb.mark_context == nil {
		linweb.mark_context = &context.Context{}
	}
	if linweb.mark_middleware == nil {
		linweb.mark_middleware = &middleware.Middleware{}
	}
	//create a new middleware to current request
	middleware := linweb.mark_middleware.New(linweb.middlewareFuncs...)
	//create a new context for current request
	context := linweb.mark_context.New(w, req, middleware)
	linweb.router.Handle(context)
}

var plugins_model interfaces.IModel

// Create a new model plugin.
func NewModel(m interface{}) interfaces.IModel {
	if plugins_model == nil {
		plugins_model = &model.Model{}
	}
	return plugins_model.New(m)
}
