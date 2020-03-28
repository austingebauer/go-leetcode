package lru_cache_146

type LRUCache struct {
	capacity int
	load     int
	keyMap   map[int]*LRUNode

	// font is always the latest used
	front *LRUNode
	// rear is always the least recently used (LRU)
	rear *LRUNode
}

type LRUNode struct {
	prev  *LRUNode
	next  *LRUNode
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keyMap:   make(map[int]*LRUNode),
	}
}

func (cache *LRUCache) Put(key int, value int) {
	existingNode, ok := cache.keyMap[key]
	if ok {
		existingNode.value = value
		cache.bringNodeToFront(existingNode)
		return
	}

	node := &LRUNode{
		key:   key,
		value: value,
	}
	cache.keyMap[key] = node

	// no need to evict, so place in front of list
	if cache.load < cache.capacity {
		if cache.front == nil && cache.rear == nil {
			cache.front = node
			cache.rear = node
		} else {
			cache.insertInFront(node)
		}

		cache.load = cache.load + 1
		return
	}

	// load is equal to capacity, so need to evict the LRU
	evicted := cache.keyMap[cache.rear.key]
	delete(cache.keyMap, cache.rear.key)

	// a single node is to be evicted
	if evicted.next == nil && evicted.prev == nil {
		cache.front = node
		cache.rear = node
		return
	}

	cache.rear = cache.rear.prev
	cache.rear.next = nil
	cache.insertInFront(node)
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.keyMap[key]
	if !ok {
		return -1
	}

	// if node is at the front of the list or is the only node in the list
	if cache.front == node || (node.prev == nil && node.next == nil) {
		return node.value
	}

	cache.bringNodeToFront(node)
	return node.value
}

func (cache *LRUCache) insertInFront(node *LRUNode) {
	cache.front.prev = node
	node.next = cache.front
	cache.front = node
}

func (cache *LRUCache) bringNodeToFront(node *LRUNode) {
	if node == cache.front {
		return
	}

	// node is the last in the list
	if node.next == nil {
		cache.rear = node.prev
	} else {
		// skip next prev to node prev
		node.next.prev = node.prev
	}

	// skip prev next to node next
	node.prev.next = node.next
	node.next = cache.front
	node.prev = nil
	cache.front.prev = node
	cache.front = node
}
