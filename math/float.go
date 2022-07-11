package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

const (
	EqualityThreshold = 1e-9
	// EqualityThreshold = math.SmallestNonzeroFloat64
)

func FloatEqual[T constraints.Float](a, b T) bool {
	return math.Abs(float64(a-b)) <= EqualityThreshold*float64(Max(1, Max(a, b)))
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
