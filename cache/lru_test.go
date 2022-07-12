package cache

import "testing"

func TestEvict(t *testing.T) {
	c := NewLRU[int, int](2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	if _, ok := c.Get(1); ok {
		t.Errorf("Should be evicted")
	}
}

func TestUpdate(t *testing.T) {
	c := NewLRU[int, int](2)
	c.Put(1, 1)
	c.Put(1, 2)
	if v, ok := c.Get(1); !ok {
		t.Errorf("Should not be evicted")
	} else if v != 2 {
		t.Errorf("Value not equal, want: %d, got: %d", 2, v)
	}
}
