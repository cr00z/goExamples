package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type proc struct {
	power  int64
	finish int64
}

type freeHeap []*proc

func (h freeHeap) Len() int           { return len(h) }
func (h freeHeap) Less(i, j int) bool { return h[i].power < h[j].power }
func (h freeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *freeHeap) Push(x interface{}) {
	*h = append(*h, x.(*proc))
}
func (h *freeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

type usedHeap []*proc

func (h usedHeap) Len() int { return len(h) }
func (h usedHeap) Less(i, j int) bool {
	if h[i].finish < h[j].finish {
		return true
	}
	if (h[i].finish == h[j].finish) && (h[i].power < h[j].power) {
		return true
	}
	return false
}
func (h usedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *usedHeap) Push(x interface{}) {
	*h = append(*h, x.(*proc))
}
func (h *usedHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

func main() {
	//now := time.Now()
	//defer func() {
	//	fmt.Println(time.Since(now))
	//}()

	var sumPower int64
	var n, m int
	var ai int64
	in := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(in, &n, &m)

	free := make(freeHeap, n)
	used := usedHeap{}

	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(in, &ai)
		free[i] = &proc{ai, 0}
	}
	heap.Init(&free)

	//for i := 0; i < 50; i++ {
	//	fmt.Println(heap.Pop(&free))
	//}

	for i := 0; i < m; i++ {
		var start, duration int64
		fmt.Fscan(in, &start, &duration)
		finish := start + duration
		for used.Len() > 0 && used[0].finish <= start {
			curr := heap.Pop(&used).(*proc)
			heap.Push(&free, curr)
		}
		if free.Len() > 0 {
			curr := heap.Pop(&free).(*proc)
			curr.finish = finish
			sumPower += curr.power * duration
			//fmt.Println("f", curr.power, duration)
			heap.Push(&used, curr)
		}
	}
	fmt.Println(sumPower)
	//fmt.Println(406407822978776)
}
