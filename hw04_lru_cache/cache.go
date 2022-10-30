package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*cacheItem
	queueMap map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*cacheItem, capacity),
		queueMap: make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	currentItem, isPresent := cache.items[key]
	switch isPresent {
	case true:
		currentItem.value = value
		cache.queue.MoveToFront(cache.queueMap[key])
	case false:
		newItem := &cacheItem{key: key, value: value}
		newListItem := cache.queue.PushFront(newItem)
		cache.items[key] = newItem
		cache.queueMap[key] = newListItem
	}

	if cache.capacity < cache.queue.Len() {
		tail := cache.queue.Back()
		cache.queue.Remove(tail)
		delete(cache.items, tail.Value.(*cacheItem).key)
		delete(cache.queueMap, tail.Value.(*cacheItem).key)
	}
	return isPresent
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	currentItem, isPresent := cache.items[key]
	if isPresent {
		cache.queue.MoveToFront(cache.queueMap[key])
		return currentItem.value, isPresent
	}
	return currentItem, isPresent
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = nil
	cache.queueMap = nil
}
