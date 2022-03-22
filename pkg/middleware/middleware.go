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

package middleware

import (
	"github.com/Codexiaoyi/linweb/interfaces"
)

var _ interfaces.IMiddleware = &Middleware{}

type Middleware struct {
	funcs []interfaces.HandlerFunc
	index int
}

func (m *Middleware) Clone() interfaces.IMiddleware {
	return &Middleware{
		funcs: make([]interfaces.HandlerFunc, 0),
		index: -1,
	}
}

func (m *Middleware) Next(context interfaces.IContext) {
	m.index++
	length := len(m.funcs)
	for ; m.index < length; m.index++ {
		m.funcs[m.index](context)
	}
}

func (m *Middleware) AddMiddlewares(middlewareFunc ...interfaces.HandlerFunc) {
	m.funcs = append(m.funcs, middlewareFunc...)
}
