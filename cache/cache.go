package cache

import "errors"

var ErrInvalidCapacity = errors.New("invalid capacity")

type Cache[K comparable, V any] interface {
	Get(K) (V, bool)
	Put(K, V)
}
