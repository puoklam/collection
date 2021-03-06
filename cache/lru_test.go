package cache

import (
	"testing"
)

func TestLRUEvict(t *testing.T) {
	c := NewLRU[int, int](2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	if _, ok := c.Get(1); ok {
		t.Errorf("Should be evicted")
	}
}

func TestLRUUpdate(t *testing.T) {
	c := NewLRU[int, int](2)
	c.Put(1, 1)
	c.Put(1, 2)
	if v, ok := c.Get(1); !ok {
		t.Errorf("Should not be evicted")
	} else if v != 2 {
		t.Errorf("Value not equal, want: %d, got: %d", 2, v)
	}
}

func TestLRUResize(t *testing.T) {
	c := NewLRU[int, int](5)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	c.Put(5, 5)
	c.Resize(3)
	tests := []struct {
		input int
		want  bool
	}{
		{input: 1, want: false},
		{input: 2, want: false},
		{input: 3, want: true},
		{input: 4, want: true},
		{input: 5, want: true},
	}
	for _, tt := range tests {
		if _, got := c.Get(tt.input); got != tt.want {
			t.Errorf("Get error after resize, want: %t, got: %t", tt.want, got)
		}
	}
}
