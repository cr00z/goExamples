package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, n, ai, first, prev, dist, max int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &k, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ai)
		if i == 0 {
			first = ai
		} else {
			dist = ai - prev
			if dist > max {
				max = dist
			}
		}
		prev = ai
	}
	dist = k - ai + first
	if dist > max {
		max = dist
	}
	fmt.Println(k - max)
}
