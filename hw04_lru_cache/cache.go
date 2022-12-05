package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       *sync.Mutex
	capacity int
	queue    List
	queueMap map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		mu:       &sync.Mutex{},
		capacity: capacity,
		queue:    NewList(),
		queueMap: make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	currentItem, isPresent := cache.queueMap[key]
	if isPresent {
		currentItem.Value.(*cacheItem).value = value
		cache.queue.MoveToFront(currentItem)
	} else {
		newItem := &cacheItem{key: key, value: value}
		newListItem := cache.queue.PushFront(newItem)
		cache.queueMap[key] = newListItem
	}

	if cache.capacity < cache.queue.Len() {
		tail := cache.queue.Back()
		cache.queue.Remove(tail)
		delete(cache.queueMap, tail.Value.(*cacheItem).key)
	}
	return isPresent
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	currentItem, isPresent := cache.queueMap[key]
	if isPresent {
		cache.queue.MoveToFront(cache.queueMap[key])
		return currentItem.Value.(*cacheItem).value, isPresent
	}
	return currentItem, isPresent
}

func (cache *lruCache) Clear() {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.queue = NewList()
	cache.queueMap = make(map[Key]*ListItem, cache.capacity)
}
