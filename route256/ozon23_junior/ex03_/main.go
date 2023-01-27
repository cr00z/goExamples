package main

import (
	"bufio"
	"fmt"
	"os"
)

type dev struct {
	pos  int
	lvl  int
	used bool
}

func absDiff(a int, b int) int {
	res := a - b
	if res < 0 {
		res = -res
	}
	return res
}

func main() {
	var t, n, ai int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]dev, n, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &ai)
			a[j] = dev{j, ai, false}
		}

		for j := 0; j < n; j++ {
			if a[j].used {
				continue
			}
			fmt.Print(j+1, " ")

			minDiff := 101
			idx := j + 1
			for k := j + 1; k < n; k++ {
				if a[k].used {
					continue
				}
				diff := absDiff(a[j].lvl, a[k].lvl)
				if diff < minDiff {
					minDiff = diff
					idx = k
				}
			}
			fmt.Println(idx + 1)
			a[idx].used = true
		}
		fmt.Println()
	}
}
