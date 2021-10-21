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
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSingletonObject(t *testing.T) {
	type test struct {
		t1 int
	}
	test1 := test{t1: 1}
	sc := newSingletonContainer()
	sc.setObject(&test1)
	getV := sc.getObject(reflect.TypeOf(&test1))
	assert.Equal(t, getV, reflect.ValueOf(&test1))
	getV1 := sc.getObject(reflect.TypeOf(&test1))
	assert.Equal(t, getV1, reflect.ValueOf(&test1))
}
