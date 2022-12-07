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
	var r, min, max int
	valid := true
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r)
		if i == 0 {
			min, max = r, r
		}
		if r > max {
			max = r
		} else if r < max {
			valid = false
		}
	}
	if valid {
		fmt.Println(max - min)
	} else {
		fmt.Println(-1)
	}
}