package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, k, u, v int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m, &k)
	destr := make(map[int]struct{}, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &u, &v)
		if u > v {
			v, u = u, v
		}
		destr[v << 16 + u] = struct{}{}
	}
	for i := 0; i
}