package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &n)
	max := 0;
	curmax := 0;
	for i := 0; i < n; i++ {
		var item int
		fmt.Fscanln(in, &item)
		if item == 1 {
			curmax++
			if curmax > max {
				max++
			}
		} else {
			curmax = 0
		}
	}
	fmt.Println(max)
}