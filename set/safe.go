package set

import (
	"fmt"
	"strings"
	"sync"
)

type safeSet[T comparable] struct {
	sync.RWMutex
	keys []T
	m    map[T]int
}

func NewSafe[T comparable](eles ...T) *safeSet[T] {
	s := &safeSet[T]{
		keys: make([]T, 0),
		m:    make(map[T]int),
	}
	s.Add(eles...)
	return s
}

func (s *safeSet[T]) Len() int {
	return len(s.m)
}

func (s *safeSet[T]) Has(eles ...T) bool {
	s.RLock()
	defer s.RUnlock()
	for _, e := range eles {
		if _, ok := s.m[e]; !ok {
			return false
		}
	}
	return true
}

func (s *safeSet[T]) Add(eles ...T) {
	s.Lock()
	defer s.Unlock()
	c := len(s.keys)
	for _, e := range eles {
		s.keys = append(s.keys, e)
		s.m[e] = c
		c++
	}
}

func (s *safeSet[T]) remove(e T) {
	if i, ok := s.m[e]; !ok {
		return
	} else {
		c := len(s.keys) - 1
		last := s.keys[c]
		s.keys[i] = last
		s.m[last] = i
		s.keys = s.keys[:c]
		delete(s.m, e)
	}
}

func (s *safeSet[T]) Remove(eles ...T) {
	s.Lock()
	defer s.Unlock()
	for _, e := range eles {
		s.remove(e)
	}
}

func (s *safeSet[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	for k := range s.m {
		delete(s.m, k)
	}
	s.keys = nil
}

func (s *safeSet[T]) IsSubset(ss Set[T]) bool {
	return isSubset[T](s, ss)
}

func (s *safeSet[T]) IsSuperset(ss Set[T]) bool {
	return ss.IsSubset(s)
}

func (s *safeSet[T]) IsIdentical(ss Set[T]) bool {
	return s.IsSubset(ss) && s.IsSuperset(ss)
}

func (s *safeSet[T]) IsDisjoint(ss Set[T]) bool {
	return isDisjoint[T](s, ss)
}

func (s *safeSet[T]) Diff(ss Set[T]) Set[T] {
	return diff[T](s, ss)
}

func (s *safeSet[T]) Items() []T {
	// return append(s.keys[:0:0], s.keys...)
	items := make([]T, len(s.keys))
	copy(items, s.keys)
	return items
}

func (s *safeSet[T]) String() string {
	s.RLock()
	defer s.RUnlock()
	eles := make([]string, 0, len(s.keys))
	for _, v := range s.keys {
		eles = append(eles, fmt.Sprintf("%v", v))
	}
	return "[" + strings.Join(eles, " ") + "]"
}
