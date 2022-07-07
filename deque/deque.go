package deque

import "container/list"

type Deque[T any] struct {
	list *list.List
}

func New[T any]() *Deque[T] {
	return &Deque[T]{
		list: list.New(),
	}
}

func (q *Deque[T]) Back() T {
	return q.list.Back().Value.(T)
}

func (q *Deque[T]) Front() T {
	return q.list.Front().Value.(T)
}

func (q *Deque[T]) Init() *Deque[T] {
	q.list.Init()
	return q
}

func (q *Deque[T]) Len() int {
	return q.list.Len()
}

func (q *Deque[T]) PushBack(v T) {
	q.list.PushBack(v)
}

func (q *Deque[T]) PushBackList(other *Deque[T]) {
	q.list.PushBackList(other.list)
}

func (q *Deque[T]) PushFront(v T) {
	q.list.PushFront(v)
}

func (q *Deque[T]) PushFrontList(other *Deque[T]) {
	q.list.PushFrontList(other.list)
}

func (q *Deque[T]) PopBack() T {
	return q.list.Remove(q.list.Back()).(T)
}

func (q *Deque[T]) PopFront() T {
	return q.list.Remove(q.list.Front()).(T)
}
