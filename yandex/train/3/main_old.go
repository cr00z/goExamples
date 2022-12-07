package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNMin(dist *[300000]int, n, k int) {
	for i := 0; i < k; i++ {
		for j := i + 1; j < n; j++ {
			if dist[j] < dist[i] {
				dist[i], dist[j] = dist[j], dist[i]
			}
		}
	}
}


func main_() {
	var n, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	var a [300000]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ready := make(map[int]int64, n)
	for i := 0; i < n; i++ {
		var sum int64
		sum, inMap := ready[a[i]]
		if !inMap {	
			var dist [300000]int
			for j := 0; j < n; j++ {
				dist[j] = a[i] - a[j]
				if dist[j] < 0 {
					dist[j] = -dist[j]
				}
			}
			//fmt.Print(dist)
			getNMin(&dist, n, k + 1)
			//fmt.Println(sortDist)
			for j := 1; j < k + 1; j++ {
				sum += int64(dist[j])
			}
			ready[a[i]] = sum
		}
		fmt.Print(sum)
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}