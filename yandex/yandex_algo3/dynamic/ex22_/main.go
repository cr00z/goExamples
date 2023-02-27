package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	if k > n {
		k = n
	}
	jumps := make([]int, n+1)
	jumps[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j > 0 {
				jumps[i] += jumps[i-j]
			} else {
				continue
			}
		}
	}
	fmt.Println(jumps[n])
}
