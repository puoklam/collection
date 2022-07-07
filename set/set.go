package set

import (
	"fmt"
	"strings"
	"sync"
)

type Set[T comparable] struct {
	sync.RWMutex
	keys []T
	m    map[T]int
}

func New[T comparable](eles ...T) *Set[T] {
	s := &Set[T]{
		keys: make([]T, 0),
		m:    make(map[T]int),
	}
	s.Add(eles...)
	return s
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Has(e T) bool {
	_, ok := s.m[e]
	return ok
}

func (s *Set[T]) Add(eles ...T) {
	s.Lock()
	defer s.Unlock()
	c := len(s.keys)
	for _, e := range eles {
		s.keys = append(s.keys, e)
		s.m[e] = c
		c++
	}
}

func (s *Set[T]) remove(e T) {
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

func (s *Set[T]) Remove(eles ...T) {
	s.Lock()
	defer s.Unlock()
	for _, e := range eles {
		s.remove(e)
	}
}

func (s *Set[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	for k := range s.m {
		delete(s.m, k)
	}
	s.keys = nil
}

func (s *Set[T]) IsSubset(ss *Set[T]) bool {
	s.RLock()
	ss.RLock()
	defer ss.RUnlock()
	defer s.RUnlock()
	for k := range s.m {
		if _, ok := ss.m[k]; !ok {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(ss *Set[T]) bool {
	return ss.IsSubset(s)
}

func (s *Set[T]) IsIdentical(ss *Set[T]) bool {
	return s.IsSubset(ss) && s.IsSuperset(ss)
}

func (s *Set[T]) IsDisjoint(ss *Set[T]) bool {
	s.RLock()
	ss.RLock()
	defer s.RUnlock()
	defer ss.RUnlock()
	for k := range s.m {
		if _, ok := ss.m[k]; ok {
			return false
		}
	}
	return true
}

func (s *Set[T]) Diff(ss *Set[T]) *Set[T] {
	s.RLock()
	ss.RLock()
	defer s.RUnlock()
	defer ss.RUnlock()

	keys := make([]T, 0)
	m := make(map[T]int)
	i := 0
	for k := range s.m {
		if _, ok := ss.m[k]; !ok {
			keys = append(keys, k)
			m[k] = i
			i++
		}
	}
	return &Set[T]{
		keys: keys,
		m:    m,
	}
}

func (s *Set[T]) String() string {
	eles := make([]string, 0, len(s.m))
	for k := range s.m {
		eles = append(eles, fmt.Sprintf("%v", k))
	}
	return "[" + strings.Join(eles, " ") + "]"
}

func Copy[T comparable](s *Set[T]) *Set[T] {
	s.RLock()
	defer s.RUnlock()
	m := make(map[T]int)
	for k, v := range s.m {
		m[k] = v
	}
	keys := make([]T, len(s.keys))
	copy(keys, s.keys)
	return &Set[T]{
		keys: keys,
		m:    m,
	}
}

func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	s := Copy(s1)
	s2.RLock()
	defer s2.RUnlock()
	c := len(s.keys)
	for k := range s2.m {
		if _, ok := s.m[k]; ok {
			continue
		}
		s.m[k] = c
		c++
	}
	return s
}

func Intersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	s1.RLock()
	s2.RLock()
	defer s1.RUnlock()
	defer s2.RUnlock()

	keys := make([]T, 0)
	m := make(map[T]int)
	i := 0
	for k := range s1.m {
		if _, ok := s2.m[k]; ok {
			keys = append(keys, k)
			m[k] = i
			i++
		}
	}
	return &Set[T]{
		keys: keys,
		m:    m,
	}
}

// https://stackoverflow.com/questions/65782592/how-to-add-a-test-device-to-firebase-cloud-messaging-fcm-in-flutter
