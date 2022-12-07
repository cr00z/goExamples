package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	pos int
	value int
}

func dist(a, b int) (dist int64) {
	if a < b {
		dist = int64(b - a)
	} else {
		dist = int64(a - b)
	}
	return
}

func main() {
	var n, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	a := make([]pair, n, n)
	for i := 0; i < n; i++ {
		a[i].pos = i
		fmt.Fscan(in, &a[i].value)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].value < a[j].value
	})

	var sums [300000]int64
	for i := 0; i < n; i++ {
		if i > 0 {
			if a[i].value == a[i-1].value {
				sums[a[i].pos] = sums[a[i-1].pos]
				continue
			}
		}

		leftPos, rightPos := i - 1, i + 1
		var leftSum, rightSum, currSum int64
		for j := 0; j < k; j++ {
			if leftPos >= 0 {
				leftSum = dist(a[leftPos].value, a[i].value)
			}
			if rightPos < n {
				rightSum = dist(a[rightPos].value, a[i].value)
			}
			if (leftPos >= 0) && (rightPos < n) {
				if leftSum < rightSum {
					currSum += leftSum
					leftPos--
				} else {
					currSum += rightSum
					rightPos++
				}
			} else if (leftPos >= 0) {
				currSum += leftSum
				leftPos--
			} else {
				currSum += rightSum
				rightPos++
			}
		}
		sums[a[i].pos] = currSum
	}

	for i := 0; i < n; i++ {
		fmt.Print(sums[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}