package cache

import (
	"container/list"
	"sync"
)

type lruNode[K comparable, V any] struct {
	k K
	v V
}

type LRUCache[K comparable, V any] struct {
	mu  sync.RWMutex
	cap int
	m   map[K]*list.Element
	l   *list.List
}

func NewLRU[K comparable, V any](cap int) *LRUCache[K, V] {
	return new(LRUCache[K, V]).Init(cap)
}

func (c *LRUCache[K, V]) add(k K, v V) {
	e := c.l.PushFront(&lruNode[K, V]{
		k: k,
		v: v,
	})
	c.m[k] = e
}

func (c *LRUCache[K, V]) evict() {
	n := c.l.Remove(c.l.Back()).(*lruNode[K, V])
	delete(c.m, n.k)
}

func (c *LRUCache[K, V]) Init(cap int) *LRUCache[K, V] {
	if cap < 1 {
		panic(ErrInvalidCapacity)
	}
	c.cap = cap
	c.m = make(map[K]*list.Element)
	c.l = list.New()
	return c
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.m[key]
	if !ok {
		var v V
		return v, false
	}
	c.l.MoveToFront(e)
	return e.Value.(*lruNode[K, V]).v, true
}

func (c *LRUCache[K, V]) Put(key K, val V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if e, ok := c.m[key]; ok {
		c.l.MoveToFront(e)
		e.Value.(*lruNode[K, V]).v = val
		return
	}
	if c.l.Len() == c.cap {
		c.evict()
	}
	c.add(key, val)
}

func (c *LRUCache[K, V]) Resize(cap int) {
	if cap < 1 {
		panic(ErrInvalidCapacity)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cap = cap
	for c.l.Len() > c.cap {
		c.evict()
	}
}

func (c *LRUCache[_, _]) Cap() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cap
}

func (c *LRUCache[_, _]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.l.Len()
}
