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

func (queue *Queue[T]) Enqueue(value T) {
	node := &Node[T]{
		Value: value,
		Next:  nil,
	}

	if queue.Length == 0 {
		queue.Tail = node
		queue.Head = node
	} else {
		queue.Tail.Next = node
		queue.Tail = node
	}

	queue.Length++
}

func (queue *Queue[T]) Dequeue() (T, bool) {
	value := queue.Head.Value
	var zeroValue T

	if queue.Length == 0 {
		return zeroValue, false
	}

	if queue.Length == 1 {
		queue.Head = nil
		queue.Tail = nil
	} else {
		queue.Head = queue.Head.Next
	}
	queue.Length--

	return value, true
}
