package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func dist(a [2]int, b[2]int) int {
	return abs(a[0],b[0]) + abs(a[1],b[1])
}

func main() {
	var n, max, start, fin int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &n)
	var cities [1000][2]int
	for i := 0; i < n; i++ {
		fmt.Fscanln(in, &cities[i][0], &cities[i][1])
	}
	fmt.Fscan(in, &max, &start, &fin)
	start--
	fin--

	var path[1000]int
	var processed[1000]bool
	for i := 0; i < n; i++ {
		path[i] = 1001
	}
	processed[start] = true

	next := 1001
	for i := 0; i < n; i++ {
		if dist(cities[i], cities[start]) <= max {
			if i == fin {
				fmt.Println(1)
				os.Exit(0)
			}
			path[i] = 1;
			next = i
		}
	}
	var min = 1;
	for i := 0; i < n; i++ {
		if processed[i] {
			continue
		}
		if next == 1001 {
			fmt.Println(-1)
			os.Exit(0)
		}
		for i := 0; i < n; i++ {
			if (dist(cities[i], cities[next]) <= max){
				if i == fin {
					fmt.Println(min)
					os.Exit(0)
				}
			    if (min + 1 < path[i]){
					path[i] = min + 1;
				}
			}
			processed[next] = true	
		}
		next_min := 1001
		for i := 0; i < n; i++ {
			if processed[i] {
				continue
			}
			if path[i] < next_min {
				next = i
				next_min = path[i]
			}
		}
		min++;
	}
	fmt.Println(-1)
}