package cache

import (
	"container/list"
	"sync"
)

type lfuNode[K comparable, V any] struct {
	c uint64
	k K
	v V
}

type LFUCache[K comparable, V any] struct {
	mu  sync.RWMutex
	cap int
	m   map[K]*list.Element
	f   map[uint64]*list.Element // heads of frequencies
	l   *list.List
}

func NewLFU[K comparable, V any](cap int) *LFUCache[K, V] {
	return new(LFUCache[K, V]).Init(cap)
}

func (c *LFUCache[K, V]) add(k K, v V) {
	n := &lfuNode[K, V]{
		c: 1,
		k: k,
		v: v,
	}
	head, ok := c.f[n.c]
	var e *list.Element
	if !ok {
		e = c.l.PushBack(n)
	} else {
		e = c.l.InsertBefore(n, head)
	}
	c.m[k] = e
	c.f[n.c] = e
}

func (c *LFUCache[K, V]) evict() {
	e := c.l.Back()
	n := c.l.Remove(e).(*lfuNode[K, V])
	delete(c.m, n.k)
	if c.f[n.c] == e {
		delete(c.f, n.c)
	}
}

func (c *LFUCache[K, V]) inc(k K) {
	e := c.m[k]
	n := e.Value.(*lfuNode[K, V])
	n.c++
	head := c.f[n.c-1]
	if head == e {
		c.f[n.c-1] = head.Next()
		if head.Next() == nil || head.Next().Value.(*lfuNode[K, V]).c != n.c-1 {
			delete(c.f, n.c-1)
		}
	}
	mark, ok := c.f[n.c]
	if !ok {
		mark = head
	}
	c.l.MoveBefore(e, mark)
	c.f[n.c] = e
}

func (c *LFUCache[K, V]) Init(cap int) *LFUCache[K, V] {
	if cap < 1 {
		panic(ErrInvalidCapacity)
	}
	c.cap = cap
	c.m = make(map[K]*list.Element)
	c.f = make(map[uint64]*list.Element)
	c.l = list.New()
	return c
}

func (c *LFUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.m[key]
	if !ok {
		var v V
		return v, false
	}
	c.inc(key)
	return e.Value.(*lfuNode[K, V]).v, true
}

func (c *LFUCache[K, V]) Put(key K, val V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.m[key]
	if ok {
		e.Value.(*lfuNode[K, V]).v = val
		c.inc(key)
		return
	}
	if c.l.Len() == c.cap {
		c.evict()
	}
	c.add(key, val)
}

func (c *LFUCache[K, V]) Resize(cap int) {
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

func (c *LFUCache[_, _]) Cap() int {
	c.mu.RLock()
	defer c.mu.RLock()
	return c.cap
}

func (c *LFUCache[_, _]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.l.Len()
}
