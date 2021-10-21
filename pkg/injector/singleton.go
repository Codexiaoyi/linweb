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

package injector

import (
	"reflect"
	"sync"
)

// The singleton object container.
type singletonContainer struct {
	singleton sync.Map
}

func newSingletonContainer() *singletonContainer {
	return &singletonContainer{
		singleton: sync.Map{},
	}
}

func (s *singletonContainer) setObject(object interface{}) {
	t := reflect.TypeOf(object)
	if t.Kind() != reflect.Ptr {
		panic(t.Name() + " not ptr type.")
	}
	s.singleton.Store(t, reflect.ValueOf(object))
}

func (s *singletonContainer) getObject(t reflect.Type) reflect.Value {
	if value, ok := s.singleton.Load(t); ok {
		return value.(reflect.Value)
	}
	res := reflect.Value{}
	//maybe object type is interface
	s.singleton.Range(func(key, value interface{}) bool {
		if key.(reflect.Type).Kind() == reflect.Interface && key.(reflect.Type).Implements(t) {
			res = value.(reflect.Value)
			return false
		}
		return true
	})
	return res
}
