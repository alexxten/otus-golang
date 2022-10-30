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
}

type cacheItem struct {
	key      Key
	value    interface{}
	listItem *ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*cacheItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	currentItem, isPresent := cache.items[key]
	switch isPresent {
	case true:
		currentItem.value = value
		cache.queue.MoveToFront(currentItem.listItem)
	case false:
		newItem := cache.queue.PushFront(value)
		cache.items[key] = &cacheItem{key: key, value: value, listItem: newItem}
	}

	if cache.capacity < cache.queue.Len() {
		tail := cache.queue.Back()
		var tailCacheItem *cacheItem
		for _, it := range cache.items {
			if it.listItem == tail {
				tailCacheItem = it
			}
		}
		cache.queue.Remove(tail)
		delete(cache.items, tailCacheItem.key)
	}
	return isPresent
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	currentItem, isPresent := cache.items[key]
	if isPresent {
		cache.queue.MoveToFront(currentItem.listItem)
		return currentItem.value, isPresent
	}
	return currentItem, isPresent
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = nil
}
