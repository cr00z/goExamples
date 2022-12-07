package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, num int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	nums := make(map[int]struct{})
	fmt.Fscan(in, &num)
	nums[num] = struct{}{}
	fmt.Println(0)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &num)
		var max int
		for j := range nums {
			cand := num ^ j
			if cand > max {
				max = cand
			}
		}
		fmt.Println(max)
		nums[num] = struct{}{}
	}
}