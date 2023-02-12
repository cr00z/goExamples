package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a, b, n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &n, &m)
	if b < a {
		a, b = b, a
		n, m = m, n
	}
	minA := n + (n-1)*a
	maxA := n + (n+1)*a
	minB := m + (m-1)*b
	maxB := m + (m+1)*b
	if minA > maxB || minB > maxA {
		fmt.Println(-1)
	} else {
		fmt.Println(max(minA, minB), min(maxA, maxB))
	}
}
