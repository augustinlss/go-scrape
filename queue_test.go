package goscrape

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	queue := NewQueue[int]()

	queue.Enqueue(10)
	if queue.Length != 1 {
		t.Errorf("Expected Length 1, got %d", queue.Length)
	}
	if queue.Head == nil || queue.Tail == nil {
		t.Fatal("Head or Tail is nil after first enqueue")
	}
	if queue.Head.Value != 10 {
		t.Errorf("Expected Head value 10, got %d", queue.Head.Value)
	}

	queue.Enqueue(20)
	if queue.Length != 2 {
		t.Errorf("Expected Length 2, got %d", queue.Length)
	}
	if queue.Tail.Value != 20 {
		t.Errorf("Expected Head value 20, got %d", queue.Head.Value)
	}
	if queue.Head.Value != 10 {
		t.Errorf("Expected Tail value 10, got %d", queue.Tail.Value)
	}
}

func TestDequeue(t *testing.T) {
	queue := NewQueue[int]()

	// Add elements

	queue.Enqueue(30)
	queue.Enqueue(20)
	queue.Enqueue(10)
	if queue.Length != 3 {
		t.Fatalf("Expected queue length 3, got %d", queue.Length)
	}

	v, _ := queue.Dequeue()
	if v != 30 {
		t.Errorf("Expected dequeued value 30, got %d", v)
	}
	if queue.Length != 2 {
		t.Errorf("Expected queue length 2, got %d", queue.Length)
	}

	v, _ = queue.Dequeue()
	if v != 20 {
		t.Errorf("Expected dequeued value 20, got %d", v)
	}

	v, _ = queue.Dequeue()
	if v != 10 {
		t.Errorf("Expected dequeued value 10, got %d", v)
	}

	if queue.Head != nil || queue.Tail != nil || queue.Length != 0 {
		t.Error("Expected empty queue after dequeuing all elements")
	}
}
