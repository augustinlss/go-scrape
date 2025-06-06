package goscrape

type Queue[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Lenght int
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Head:   nil,
		Tail:   nil,
		Lenght: 0,
	}
}
