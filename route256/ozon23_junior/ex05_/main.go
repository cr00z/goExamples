package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, ai, prev int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(in, &n)
		a := make(map[int]struct{}, n)
		crit := "YES"
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &ai)
			if i != 0 && ai != prev {
				if _, inMap := a[ai]; inMap {
					crit = "NO"
				}
			}
			a[ai] = struct{}{}
			prev = ai
		}
		fmt.Println(crit)
	}
}
