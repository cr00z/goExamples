package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(x int64, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b, c, d int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c, &d)
	minm := min(a, b)
	minn := min(c, d)
	fmt.Println(minm + 1, minn + 1)
}