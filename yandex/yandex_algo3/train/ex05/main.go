package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	let := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &let[i])
	}
	result := 0
	for i := 0; i < n; i++ {
		cnt := let[i]
		rng := 0
		for j := i + 1; j < n && let[j] != 0; j++ {
			if let[i] < cnt {
				cnt = let[i]
			}
			rng++
		}
		for j := i + 1; j < i+1+rng; j++ {
			let[j] -= cnt
		}
		result += rng * cnt
	}
	fmt.Println(result)
}
