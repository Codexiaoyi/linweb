//Copyright 2021 Codexiaoyi
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package linweb

import (
	"net/http"
	"sync"

	"github.com/Codexiaoyi/linweb/interfaces"
	"github.com/Codexiaoyi/linweb/pkg/model"
)

var (
	pluginsModel interfaces.IModel
	Cache        interfaces.ICache
)

type Linweb struct {
	markRouter  interfaces.IRouter
	markContext interfaces.IContext
	// this middleware means an implement, every request need create a new middleware from New() by markMiddleware.
	markMiddleware interfaces.IMiddleware
	middlewareFunc []interfaces.HandlerFunc
	markModel      interfaces.IModel
	markInject     interfaces.IInjector
	markCache      interfaces.ICache

	contextPool sync.Pool
}

// Create a new Linweb.
// Add customize plugins with method of plugins.go, otherwise use default plugins.
func NewLinweb(plugins ...CustomizePlugins) *Linweb {
	lin := &Linweb{}
	defaultPlugins()(lin)
	for _, plugin := range plugins {
		plugin(lin)
	}
	pluginsModel = lin.markModel
	Cache = lin.markCache
	return lin
}

// AddSingleton Add objects to DI container with a single instance in every request, they must all be of pointer type.
func (lin *Linweb) AddSingleton(objs ...interface{}) {
	lin.markInject.AddSingleton(objs...)
}

// AddTransient Add objects to DI container with new instance in every request, they must all be of pointer type.
func (lin *Linweb) AddTransient(objs ...interface{}) {
	lin.markInject.AddTransient(objs...)
}

// AddControllers Add all controllers, they must all be of pointer type
func (lin *Linweb) AddControllers(obj ...interface{}) {
	lin.markRouter.AddControllers(obj)
}

// AddMiddlewares Add global middlewares
func (lin *Linweb) AddMiddlewares(middlewareFunc ...interfaces.HandlerFunc) {
	lin.middlewareFunc = middlewareFunc
}

// Run you project to listen the "addr", enjoy yourself!
func (lin *Linweb) Run(addr string) error {
	lin.contextPool.New = func() interface{} {
		//clone a new context for current request
		return lin.markContext.Clone()
	}
	return http.ListenAndServe(addr, lin)
}

// Serve HTTP.
func (lin *Linweb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//clone a new middleware to current request
	middleware := lin.markMiddleware.Clone()
	//add user middlewares
	middleware.AddMiddlewares(lin.middlewareFunc...)
	// get a context from pool
	ctx := lin.contextPool.Get().(interfaces.IContext)
	ctx.Reset(w, req, middleware)

	lin.markRouter.Handle(ctx, lin.markInject)

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
