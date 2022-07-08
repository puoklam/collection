package set

import (
	"errors"
	"sync"
)

var ErrTypeNotMatch = errors.New("Type not match")

type locker = sync.Locker
type rwLocker interface {
	RLock()
	RUnlock()
}

type Set[T comparable] interface {
	Len() int
	Has(eles ...T) bool
	Add(eles ...T)
	Remove(eles ...T)
	Clear()
	Items() []T
	Diff(ss Set[T]) Set[T]
	IsSubset(ss Set[T]) bool
	IsSuperset(ss Set[T]) bool
	IsIdentical(ss Set[T]) bool
	IsDisjoint(ss Set[T]) bool
}

func diff[T comparable](s1, s2 Set[T]) Set[T] {
	if l, ok := s1.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	if l, ok := s2.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}

	keys := make([]T, 0)
	m := make(map[T]int)
	i := 0
	for _, v := range s1.Items() {
		if !s2.Has(v) {
			keys = append(keys, v)
			m[v] = i
			i++
		}
	}
	return &safeSet[T]{
		keys: keys,
		m:    m,
	}
}

func isSubset[T comparable](s1, s2 Set[T]) bool {
	if l, ok := s1.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	if l, ok := s2.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	for _, v := range s1.Items() {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

func isDisjoint[T comparable](s1, s2 Set[T]) bool {
	if l, ok := s1.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	if l, ok := s2.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	for _, v := range s1.Items() {
		if s2.Has(v) {
			return false
		}
	}
	return true
}

func Copy[T comparable](s Set[T]) Set[T] {
	if l, ok := s.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	m := make(map[T]int)
	for i, v := range s.Items() {
		m[v] = i
	}
	keys := make([]T, s.Len())
	copy(keys, s.Items())
	return &safeSet[T]{
		keys: keys,
		m:    m,
	}
}

func Union[T comparable](sets ...Set[T]) Set[T] {
	for _, s := range sets {
		if l, ok := s.(locker); ok {
			l.Lock()
			defer l.Unlock()
		}
	}
	keys := make([]T, 0)
	m := make(map[T]int)
	i := 0
	for _, s := range sets {
		for _, v := range s.Items() {
			if _, ok := m[v]; !ok {
				keys = append(keys, v)
				m[v] = i
				i++
			}
		}
	}
	return &safeSet[T]{
		keys: keys,
		m:    m,
	}
}

func Intersection[T comparable](s1, s2 Set[T]) Set[T] {
	if l, ok := s1.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	if l, ok := s2.(rwLocker); ok {
		l.RLock()
		defer l.RUnlock()
	}
	keys := make([]T, 0)
	m := make(map[T]int)
	i := 0
	for _, v := range s1.Items() {
		if s2.Has(v) {
			keys = append(keys, v)
			m[v] = i
			i++
		}
	}
	return &safeSet[T]{
		keys: keys,
		m:    m,
	}
}
