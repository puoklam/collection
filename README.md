# collection
[![GoDoc Widget]][GoDoc]
Collections of data structures and utilities written in Go

## Install

`go get github.com/pukolam/collection`


## Usage

### Set

```go
s1, s2 := set.NewSafe(1, 2), set.NewSafe(4, 5)
s1.Add(3)
s1.Remove(2)
u := set.Union[int](s1, s2)
fmt.Println(u) // [1 ,3 ,4 ,5]
```

### Deque

```go
q := deque.New[int]()
q.PushBack(1)
q.PushBack(2)
q.PushBack(3)
fmt.Println(q.PopFront()) // 1
fmt.Println(q.PopBack()) // 3
```