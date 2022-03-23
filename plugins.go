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
	"github.com/Codexiaoyi/linweb/interfaces"
	"github.com/Codexiaoyi/linweb/pkg/cache"
	"github.com/Codexiaoyi/linweb/pkg/context"
	"github.com/Codexiaoyi/linweb/pkg/injector"
	"github.com/Codexiaoyi/linweb/pkg/middleware"
	"github.com/Codexiaoyi/linweb/pkg/model"
	"github.com/Codexiaoyi/linweb/pkg/router"
)

// CustomizePlugins Add customize plugins.
// It is not allowed to pass in non-plugin implementations.
// Without customize plugins will use the default plugins.
type CustomizePlugins func(lin *Linweb)

func defaultPlugins() CustomizePlugins {
	return func(lin *Linweb) {
		lin.markRouter = router.New()
		lin.markContext = &context.Context{}
		lin.markMiddleware = &middleware.Middleware{}
		lin.markInject = injector.Instance()
		lin.markCache = cache.Instance()
		lin.markModel = &model.Model{}
	}
}

// Customize router plugin.
func RouterPlugin(router interfaces.IRouter) CustomizePlugins {
	return func(lin *Linweb) {
		lin.markRouter = router
	}
}

// Customize context plugin.
func ContextPlugin(context interfaces.IContext) CustomizePlugins {
	return func(lin *Linweb) {
		lin.markContext = context
	}
}

// Customize middleware plugin.
func MiddlewarePlugin(middleware interfaces.IMiddleware) CustomizePlugins {
	return func(lin *Linweb) {
		lin.markMiddleware = middleware
	}
}

// Customize inject plugin.
func InjectPlugin(inject interfaces.IInjector) CustomizePlugins {
	return func(lin *Linweb) {
		lin.markInject = inject
	}
}

// Customize cache plugin.
func CachePlugin(cache interfaces.ICache) CustomizePlugins {
	return func(lin *Linweb) {
		lin.markCache = cache
	}
}

// Customize model plugin.
func ModelPlugin(model interfaces.IModel) CustomizePlugins {
	return func(lin *Linweb) {
		lin.markModel = model
	}
}
