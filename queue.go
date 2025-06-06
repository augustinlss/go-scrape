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

func Enqueue[T any](queue *Queue[T], value T) {
	node := Node[T]{
		Value: value,
		Next:  queue.Head,
	}

	if queue.Lenght == 0 {
		queue.Tail = &node
		queue.Head = &node
	} else {
		queue.Head = &node
	}

	queue.Lenght++
}
