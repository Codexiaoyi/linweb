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
	"container/list"
	"sync"
	"unsafe"
)

type lru struct {
	// the number of visits less than k in the history list.
	k int8
	// already all used bytes, only computes size of keys and values.
	currentBytes int64
	// doubly link list.
	cacheList *list.List
	// history link list.
	historyList *list.List
	// save key and element of linkList position.
	//cacheMap map[string]*list.Element
	cacheMap sync.Map
	// callback when the key deleted.
	onLruDelete func(key string)
}

type node struct {
	onHistory bool
	count     int8
	key       string
	value     interface{}
}

func newLru(k int8, onLruDelete func(key string)) *lru {
	return &lru{
		k:           k,
		cacheList:   list.New(),
		historyList: list.New(),
		onLruDelete: onLruDelete,
	}
}

func (l *lru) add(key string, value interface{}) {
	// key is exist, update value
	if ele, ok := l.cacheMap.Load(key); ok {
		element := ele.(*list.Element)
		node := element.Value.(*node)
		l.currentBytes += int64(unsafe.Sizeof(value)) - int64(unsafe.Sizeof(node.value))
		node.value = value
		if node.onHistory {
			node.count++
			if node.count >= l.k {
				l.moveToCache(element)
			}
		}
	} else {
		// new cache insert to history list.
		element := l.historyList.PushBack(&node{
			onHistory: true,
			count:     0,
			key:       key,
			value:     value,
		})
		l.cacheMap.Store(key, element)
		l.currentBytes += int64(unsafe.Sizeof(key)) + int64(unsafe.Sizeof(value))
	}
}

func (l *lru) get(key string) (interface{}, bool) {
	if ele, ok := l.cacheMap.Load(key); ok {
		element := ele.(*list.Element)
		node := element.Value.(*node)
		if node.onHistory {
			node.count++
			if node.count >= l.k {
				l.moveToCache(element)
			}
		} else {
			l.cacheList.MoveToBack(element)
		}
		return node.value, true
	}
	return nil, false
}

func (l *lru) delete(key string) {
	if ele, ok := l.cacheMap.Load(key); ok {
		element := ele.(*list.Element)
		node := element.Value.(*node)
		if node.onHistory {
			l.historyList.Remove(element)
		} else {
			l.cacheList.Remove(element)
		}
		l.cacheMap.Delete(node.key)
		l.currentBytes -= int64(len(node.key)) + int64(unsafe.Sizeof(node.value))
		// delete completed, notice deleted event.
		if l.onLruDelete != nil {
			l.onLruDelete(node.key)
		}
	}
}

// Remove the oldest element in front of cache list.
func (l *lru) removeOldest() {
	element := l.cacheList.Front()
	if element != nil {
		node := element.Value.(*node)
		l.delete(node.key)
	}
}

// Move element to cache list from history list.
func (l *lru) moveToCache(element *list.Element) {
	l.historyList.Remove(element)
	node := element.Value.(*node)
	node.onHistory = false
	ele := l.cacheList.PushBack(node)
	l.cacheMap.Store(node.key, ele)
}
