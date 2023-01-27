package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, pi int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		p := make(map[int]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &pi)
			p[pi]++
		}
		sum := 0
		for k, v := range p {
			sum += k * (v/3*2 + v%3)
		}
		fmt.Println(sum)
	}
}
