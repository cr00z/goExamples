package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(brev []int, maxBrev int, k int) bool {
	sum := 0
	for i := range brev {
		if brev[i] > maxBrev {
			sum += (brev[i] - 1) / maxBrev
		}
	}
	return sum <= k
}

func main() {
	var n, k, last int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	brev := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &brev[i])
		if brev[i] > last {
			last = brev[i]
		}
	}
	first := last / 2
	for first != last {
		if check(brev, first, k) {
			last = first
			first = last / 2
		} else {
			if first == last-1 {
				break
			}
			first = (last + first) / 2
		}
	}
	fmt.Println(last)
}
