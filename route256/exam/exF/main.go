package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var n2 int
		fmt.Fscan(in, &n2)
		queue := list.New()
		for j := 0; j < n2; j++ {
			item := [3]int{}
			fmt.Fscan(in, &item[0], &item[1], &item[2])
			queue.PushBack(item)
		}
		buf := make([]int, n2, n2)
		e := queue.Front() // Первый элемент
		buf[0] = e.Value.([3]int)[0]
		buf[1] = e.Value.([3]int)[1]
		i := 2
		queue.Remove(e)
		for queue.Len() > 0 {
			e := queue.Front()
			item := e.Value.([3]int)
			queue.Remove(e)
			if buf[i-1] == item[0] {
				if buf[i-2] != item[1] {
					buf[i] = item[1]
				} else {
					buf[i] = item[2]
				}
				i++
				if i == n2 {
					break
				}
			} else {
				queue.PushBack(item)
			}
		}
		for j := 0; j < n2/2; j++ {
			fmt.Println(buf[j], buf[n2/2+j])
		}
		fmt.Println()
	}
}
