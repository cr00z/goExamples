package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c)
	if c < 0 {
		fmt.Println("NO SOLUTION")
	} else if (a == 0) && (b > 0) && (b*b == c) {
		fmt.Println("MANY SOLUTIONS")
	} else {
		if a != 0 {
			x := (c*c - b) / a
			x2 := float64(c*c-b) / float64(a)
			if float64(x) == x2 {
				fmt.Println(x)
			} else {
				fmt.Println("NO SOLUTION")
			}
		} else {
			fmt.Println("NO SOLUTION")
		}
	}
}
