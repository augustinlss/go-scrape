package goscrape

type Queue[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func Enqueue[T any](queue *Queue[T], value T) {
	node := &Node[T]{
		Value: value,
		Next:  queue.Head,
	}

	if queue.Length == 0 {
		queue.Tail = node
		queue.Head = node
	} else {
		queue.Head = node
	}

	queue.Length++
}

func Dequeue[T any](queue *Queue[T]) T {
	value := queue.Head.Value
	if queue.Length == 1 {
		queue.Head = nil
		queue.Tail = nil
	} else {
		queue.Head = queue.Head.Next
	}
	queue.Length--

	return value
}
