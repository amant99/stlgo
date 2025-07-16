// ---- main.go (final demo usage of Queue, Stack, PriorityQueue) ----
package main

import (
	"fmt"

	"github.com/amant99/stlgo/internal/factory"
)

func main() {
	// --- Queue Example ---
	fmt.Println("--- Queue Demo ---")
	q := factory.NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println("Queue Size:", q.Size())
	val, ok := q.Dequeue()
	if ok {
		fmt.Println("Dequeued:", val)
	}
	peek, _ := q.Peek()
	fmt.Println("Next in Queue:", peek)

	// --- Stack Example ---
	fmt.Println("\n--- Stack Demo ---")
	s := factory.NewStack[string]()
	s.Push("hello")
	s.Push("world")
	fmt.Println("Stack Size:", s.Size())
	sval, _ := s.Pop()
	fmt.Println("Popped from stack:", sval)
	speek, _ := s.Peek()
	fmt.Println("Top of Stack:", speek)

	// --- Priority Queue Example (min-heap) ---
	fmt.Println("\n--- Priority Queue Demo (Min-Heap) ---")
	pq := factory.NewPriorityQueue[int](func(a, b int) bool { return a < b })
	pq.PushValue(5)
	pq.PushValue(1)
	pq.PushValue(3)
	fmt.Println("Priority Queue Size:", pq.Size())
	pval, _ := pq.PopValue()
	fmt.Println("Popped from priority queue:", pval)
	ppeek, _ := pq.Peek()
	fmt.Println("Next in Priority Queue:", ppeek)
}
