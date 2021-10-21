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

func TestGetNewObj(t *testing.T) {
	type test struct {
		t1 int
	}
	test1 := test{t1: 1}
	tc := newTransientContainer()

	test2 := tc.getNewObj(reflect.ValueOf(&test1)).Interface().(*test)
	test2.t1 = 2

	assert.NotEqual(t, test1, test2)
	assert.Equal(t, test1.t1, 1)
	assert.Equal(t, test2.t1, 2)
}
