package set

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type unsafeSet[T comparable] struct {
	keys []T
	m    map[T]int
	r    *rand.Rand
}

func NewUnSafe[T comparable](eles ...T) *unsafeSet[T] {
	s := &unsafeSet[T]{
		keys: make([]T, 0),
		m:    make(map[T]int),
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	s.Add(eles...)
	return s
}

func (s *unsafeSet[T]) Len() int {
	return len(s.m)
}

func (s *unsafeSet[T]) Has(eles ...T) bool {
	for _, e := range eles {
		if _, ok := s.m[e]; !ok {
			return false
		}
	}
	return true
}

func (s *unsafeSet[T]) Add(eles ...T) {
	c := len(s.keys)
	for _, e := range eles {
		s.keys = append(s.keys, e)
		s.m[e] = c
		c++
	}
}

func (s *unsafeSet[T]) remove(e T) {
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

func (s *unsafeSet[T]) Remove(eles ...T) {
	for _, e := range eles {
		s.remove(e)
	}
}

func (s *unsafeSet[T]) Clear() {
	for k := range s.m {
		delete(s.m, k)
	}
	s.keys = nil
}

func (s *unsafeSet[T]) IsSubset(ss Interface[T]) bool {
	return isSubset[T](s, ss)
}

func (s *unsafeSet[T]) IsSuperset(ss Interface[T]) bool {
	return ss.IsSubset(s)
}

func (s *unsafeSet[T]) IsIdentical(ss Interface[T]) bool {
	return s.IsSubset(ss) && s.IsSuperset(ss)
}

func (s *unsafeSet[T]) IsDisjoint(ss Interface[T]) bool {
	return isDisjoint[T](s, ss)
}

func (s *unsafeSet[T]) Diff(ss Interface[T]) Interface[T] {
	return diff[T](s, ss)
}

func (s *unsafeSet[T]) Items() []T {
	items := make([]T, len(s.keys))
	copy(items, s.keys)
	return items
}

func (s *unsafeSet[T]) String() string {
	eles := make([]string, 0, len(s.keys))
	for _, v := range s.keys {
		eles = append(eles, fmt.Sprintf("%v", v))
	}
	return "[" + strings.Join(eles, " ") + "]"
}

func (s *unsafeSet[T]) Seed(seed int64) {
	s.r.Seed(seed)
}

func (s *unsafeSet[T]) Random() T {
	return s.keys[s.r.Intn(s.Len())]
}
