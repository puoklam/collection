package cache

import (
	"sync"
)

type node[K comparable, V any] struct {
	prev, next *node[K, V]
	k          K
	v          V
}

type LRUCache[K comparable, V any] struct {
	mu   sync.RWMutex
	cap  int
	len  int
	root *node[K, V]
	m    map[K]*node[K, V]
}

func NewLRU[K comparable, V any](capacity int) *LRUCache[K, V] {
	return new(LRUCache[K, V]).Init(capacity)
}

func (c *LRUCache[K, V]) add(n *node[K, V]) *node[K, V] {
	r := c.root
	n.prev = r
	n.next = r.next
	n.prev.next = n
	n.next.prev = n
	c.len++
	return n
}

func (c *LRUCache[K, V]) addPair(key K, val V) *node[K, V] {
	return c.add(&node[K, V]{
		k: key,
		v: val,
	})
}

func (c *LRUCache[K, V]) evict() *node[K, V] {
	r := c.root
	l := r.prev
	l.prev.next = l.next
	l.next.prev = l.prev
	if l != r {
		c.len--
		l.prev, l.next = nil, nil
	}
	return l
}

func (c *LRUCache[K, V]) refresh(n *node[K, V]) {
	r := c.root
	if n == r.next {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = r
	n.next = r.next
	n.prev.next = n
	n.next.prev = n
}

func (c *LRUCache[K, V]) Init(capacity int) *LRUCache[K, V] {
	if capacity < 1 {
		panic("invalid capacity")
	}
	c.cap = capacity
	c.len = 0
	c.m = make(map[K]*node[K, V])
	c.root = &node[K, V]{}
	c.root.prev = c.root
	c.root.next = c.root
	return c
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	n, ok := c.m[key]
	if !ok {
		var v V
		return v, false
	}
	c.refresh(n)
	return n.v, true
}

func (c *LRUCache[K, V]) Put(key K, val V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if n, ok := c.m[key]; ok {
		c.refresh(n)
		n.v = val
		return
	}
	if c.len == c.cap {
		l := c.evict()
		delete(c.m, l.k)
	}
	n := c.addPair(key, val)
	c.m[key] = n
}

func (c *LRUCache[_, _]) Cap() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cap
}

func (c *LRUCache[_, _]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.len
}
