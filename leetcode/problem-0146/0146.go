package problem

import "container/list"

type kv struct {
	k int
	v int
}

// LRUCache ...
type LRUCache struct {
	capacity int
	data     map[int]*list.Element
	hits     *list.List
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*list.Element),
		hits:     list.New(),
	}
}

// Get ...
func (c *LRUCache) Get(key int) int {
	if ele, ok := c.data[key]; ok {
		c.hits.MoveToFront(ele)
		return ele.Value.(kv).v
	}

	return -1
}

// Put ...
func (c *LRUCache) Put(key int, value int) {
	if ele, ok := c.data[key]; ok {
		ele.Value = kv{k: key, v: value}
		c.hits.MoveToFront(ele)
		return
	}

	if c.capacity == len(c.data) {
		last := c.hits.Back()
		delete(c.data, last.Value.(kv).k)
		c.hits.Remove(last)
	}

	c.data[key] = c.hits.PushFront(kv{k: key, v: value})
}
