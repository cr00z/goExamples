package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type op struct {
	step int
	prev int
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	dp := make([]op, n+1)
	for i := 2; i <= n; i++ {
		add := dp[i-1].step
		prev := i - 1
		if i%2 == 0 {
			if dp[i/2].step < add {
				add = dp[i/2].step
				prev = i / 2
			}
		}
		if i%3 == 0 {
			if dp[i/3].step < add {
				add = dp[i/3].step
				prev = i / 3
			}
		}
		dp[i].step = add + 1
		dp[i].prev = prev
	}

	fmt.Println(dp[n].step)
	steps := make([]int, 0)
	for n != 0 {
		steps = append(steps, n)
		n = dp[n].prev
	}
	sort.Ints(steps)
	fmt.Println(strings.Trim(fmt.Sprint(steps), "[]"))
}
