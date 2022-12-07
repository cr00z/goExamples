package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, cnt, candidate int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	var prized [100001]bool
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a)
		if (a == i) || prized[a] {
			cnt++
			if cnt == 1 {
				candidate = i
			}
		} else {
			prized[a] = true
		}
	}
	if cnt != 1 {
		fmt.Println(-1, -1)
	} else {
		for i := 1; i <= n; i++ {
			if prized[i] == false {
				fmt.Println(candidate, i)
				break
			}
		}
	}
}