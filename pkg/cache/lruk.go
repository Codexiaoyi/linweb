package cache

import (
	"container/list"
	"linweb/interfaces"
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
	cacheMap map[string]*list.Element
}

type node struct {
	onHistory bool
	count     int8
	key       string
	value     interfaces.Value
}

func newLru(k int8) *lru {
	return &lru{
		k:           k,
		cacheList:   list.New(),
		historyList: list.New(),
		cacheMap:    make(map[string]*list.Element),
	}
}

func (l *lru) add(key string, value interfaces.Value) {
	// key is exist, update value
	if element, ok := l.cacheMap[key]; ok {
		node := element.Value.(*node)
		l.currentBytes += int64(value.Len()) - int64(node.value.Len())
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
		l.cacheMap[key] = element
		l.currentBytes += int64(len(key)) + int64(value.Len())
	}
}

func (l *lru) get(key string) (interfaces.Value, bool) {
	if element, ok := l.cacheMap[key]; ok {
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
	if element, ok := l.cacheMap[key]; ok {
		node := element.Value.(*node)
		if node.onHistory {
			l.historyList.Remove(element)
		} else {
			l.cacheList.Remove(element)
		}
		delete(l.cacheMap, node.key)
		l.currentBytes -= int64(len(node.key)) + int64(node.value.Len())
	}
}

// Remove the oldest element in front of cache list.
func (l *lru) removeOldest() {
	element := l.cacheList.Front()
	if element != nil {
		l.cacheList.Remove(element)
		node := element.Value.(*node)
		delete(l.cacheMap, node.key)
		l.currentBytes -= int64(len(node.key)) + int64(node.value.Len())
	}
}

// Move element to cache list from history list.
func (l *lru) moveToCache(element *list.Element) {
	l.historyList.Remove(element)
	node := element.Value.(*node)
	node.onHistory = false
	ele := l.cacheList.PushBack(node)
	l.cacheMap[node.key] = ele
}
