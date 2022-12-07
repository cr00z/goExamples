package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, first, second, cnt, all1, all2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a)
		if (a % 2 == 0) {
			cnt++
			if (i % 2 == 1) {
				first = i
				all1++
			}
		} else {
			if (i % 2 == 0) {
				second = i
				all2++
			}
		}
		
	}
	if (all1 == 1) && (all2 == 1) && (cnt == n / 2) {
		if first < second {
			fmt.Println(first, second)
		} else {
			fmt.Println(second, first)
		}		
	} else {
		fmt.Println(-1, -1)
	}
}