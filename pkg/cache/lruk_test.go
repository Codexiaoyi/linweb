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

package cache

import (
	"testing"
	"unsafe"

	"gopkg.in/go-playground/assert.v1"
)

type TestValue struct {
	v string
}

func TestAdd(t *testing.T) {
	lru := newLru(2, func(key string) {
		println(key)
	})
	lru.add("key1", &TestValue{v: "test string...."})
	assert.Equal(t, lru.historyList.Len(), 1)
	assert.Equal(t, lru.cacheList.Len(), 0)
	assert.Equal(t, lru.currentBytes, int64(unsafe.Sizeof("key1")+unsafe.Sizeof("test string....")))

	lru.add("key1", &TestValue{v: "cover"})
	assert.Equal(t, lru.historyList.Len(), 1)
	assert.Equal(t, lru.cacheList.Len(), 0)
	assert.Equal(t, lru.currentBytes, int64(unsafe.Sizeof("key1")+unsafe.Sizeof("cover")))

	lru.add("key1", &TestValue{v: "cover"})
	assert.Equal(t, lru.historyList.Len(), 0)
	assert.Equal(t, lru.cacheList.Len(), 1)
}

func TestGet(t *testing.T) {
	lru := newLru(2, func(key string) {
		println(key)
	})
	lru.add("key1", &TestValue{v: "test string...."})
	lru.get("key1")
	assert.Equal(t, lru.historyList.Len(), 1)
	assert.Equal(t, lru.cacheList.Len(), 0)
	lru.get("key1")
	assert.Equal(t, lru.historyList.Len(), 0)
	assert.Equal(t, lru.cacheList.Len(), 1)

	lru.add("key2", &TestValue{v: "test string...."})
	lru.get("key2")
	lru.get("key2")
	assert.Equal(t, lru.historyList.Len(), 0)
	assert.Equal(t, lru.cacheList.Len(), 2)

	assert.Equal(t, lru.cacheList.Back().Value.(*node).key, "key2")
	lru.get("key1")
	assert.Equal(t, lru.cacheList.Back().Value.(*node).key, "key1")
}
