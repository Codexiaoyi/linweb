package linweb

import (
	"fmt"
	"linweb/interfaces"
	"linweb/pkg/cache"
	"linweb/pkg/context"
	"linweb/pkg/middleware"
	"linweb/pkg/model"
	"linweb/pkg/router"
	"log"
	"net/http"
	"reflect"
	"sync"
)

var (
	pluginsModel interfaces.IModel
	Cache        interfaces.ICache
)

type LinWeb struct {
	router      interfaces.IRouter
	markContext interfaces.IContext
	// this middleware means an implement, every request need create a new middleware from New() by markMiddleware.
	markMiddleware interfaces.IMiddleware
	middlewareFunc []interfaces.HandlerFunc

	contextPool sync.Pool
}

func NewLinWeb() *LinWeb {
	return &LinWeb{}
}

// AddCustomizePlugins Add customize plugins, they must all be of pointer type.
// It is not allowed to pass in non-plugin implementations.
// Without customize plugins will use the default plugins.
func (lin *LinWeb) AddCustomizePlugins(plugins ...interface{}) {
	for _, p := range plugins {
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.ICache)(nil)).Elem()) {
			Cache = p.(interfaces.ICache)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IRouter)(nil)).Elem()) {
			lin.router = p.(interfaces.IRouter)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IContext)(nil)).Elem()) {
			lin.markContext = p.(interfaces.IContext)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IModel)(nil)).Elem()) {
			pluginsModel = p.(interfaces.IModel)
			continue
		}
		if reflect.TypeOf(p).Implements(reflect.TypeOf((*interfaces.IMiddleware)(nil)).Elem()) {
			lin.markMiddleware = p.(interfaces.IMiddleware)
			continue
		}
		log.Fatal(fmt.Sprintf("'%s' is not a plugin, please check if it implements a plugin", reflect.TypeOf(p).Elem().Name()))
	}
}

// AddControllers Add all controllers, they must all be of pointer type
func (lin *LinWeb) AddControllers(obj ...interface{}) {
	if lin.router == nil {
		lin.router = router.New()
	}
	lin.router.AddControllers(obj)
}

// AddMiddlewares Add global middlewares
func (lin *LinWeb) AddMiddlewares(middlewareFunc ...interfaces.HandlerFunc) {
	lin.middlewareFunc = middlewareFunc
}

// Run you project to listen the "addr", enjoy yourself!
func (lin *LinWeb) Run(addr string) error {
	if lin.markContext == nil {
		lin.markContext = &context.Context{}
	}
	if lin.markMiddleware == nil {
		lin.markMiddleware = &middleware.Middleware{}
	}
	if Cache == nil {
		Cache = cache.Instance()
	}
	lin.contextPool.New = func() interface{} {
		//create a new context for current request
		return lin.markContext.New()
	}
	return http.ListenAndServe(addr, lin)
}

func (lin *LinWeb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//create a new middleware to current request
	middleware := lin.markMiddleware.New(lin.middlewareFunc...)
	// get a context from pool
	ctx := lin.contextPool.Get().(interfaces.IContext)
	ctx.Reset(w, req, middleware)

	lin.router.Handle(ctx)

	//end handle, take context to pool
	lin.contextPool.Put(ctx)
}

// NewModel Create a new model plugin.
func NewModel(m interface{}) interfaces.IModel {
	if pluginsModel == nil {
		pluginsModel = &model.Model{}
	}
	return pluginsModel.New(m)
}
