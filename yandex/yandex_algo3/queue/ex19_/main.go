package main

import (
	"bufio"
	"fmt"
	"os"
)

func leftChildIdx(i int) int {
	return 2*i + 1
}

func rightChildIdx(i int) int {
	return 2*i + 2
}

func parentIdx(i int) int {
	return (i - 1) / 2
}

type Heap[T any] struct {
	data []T
	comp func(a, b T) bool
}

func NewHeap[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		comp: comp,
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		var val T
		return val, false
	}
	return h.data[0], true
}

func (h *Heap[T]) heapifyUp(i int) {
	for h.comp(h.data[parentIdx(i)], h.data[i]) {
		h.swap(i, parentIdx(i))
		i = parentIdx(i)
	}
}

func (h *Heap[T]) heapifyDown(i int) {
	l, r := leftChildIdx(i), rightChildIdx(i)
	largest := i
	if l < h.Size() && h.comp(h.data[i], h.data[l]) {
		largest = l
	}
	if r < h.Size() && h.comp(h.data[largest], h.data[r]) {
		largest = r
	}
	if largest != i {
		h.swap(i, largest)
		h.heapifyDown(largest)
	}
}

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.heapifyUp(h.Size() - 1)
}

func (h *Heap[T]) Pop() (T, bool) {
	var val T
	if h.Size() == 0 {
		return val, false
	}

	val = h.data[0]
	h.swap(0, h.Size()-1)
	h.data = h.data[:h.Size()-1]
	h.heapifyDown(0)

	return val, true
}

func main() {
	heap := NewHeap(func(a, b int) bool { return a < b })

	var n, cmd, op int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &cmd)
		if cmd == 0 {
			fmt.Fscan(in, &op)
			heap.Push(op)
		} else {
			num, _ := heap.Pop()
			fmt.Println(num)
		}
	}
}
