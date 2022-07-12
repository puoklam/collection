package cache

import "container/list"

type node[K comparable, V any] struct {
	k K
	v V
}

type LRUCache[K comparable, V any] struct {
	c uint32
	m map[K]*list.Element
	l *list.List
}

func NewLRU[K comparable, V any](capacity uint32) *LRUCache[K, V] {
	if capacity == 0 {
		panic("invalid capacity")
	}
	return &LRUCache[K, V]{
		c: capacity,
		m: make(map[K]*list.Element),
		l: list.New(),
	}
}

func (c *LRUCache[K, V]) remove(key K) {
	e, ok := c.m[key]
	if !ok {
		return
	}
	c.l.Remove(e)
	delete(c.m, key)

}

func (c *LRUCache[K, V]) evict() {
	n := c.l.Remove(c.l.Back()).(*node[K, V])
	delete(c.m, n.k)
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	e, ok := c.m[key]
	if !ok {
		var v V
		return v, false
	}
	c.l.MoveToFront(e)
	return e.Value.(*node[K, V]).v, true
}

func (c *LRUCache[K, V]) Put(key K, val V) {
	c.remove(key)
	if uint32(len(c.m)) == c.c {
		c.evict()
	}
	e := c.l.PushFront(&node[K, V]{
		k: key,
		v: val,
	})
	c.m[key] = e
}

func (c *LRUCache[_, _]) Cap() uint32 {
	return c.c
}

func (c *LRUCache[_, _]) Len() int {
	return c.l.Len()
}
