package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c float64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c)
	if a == 0 && b == 0 && c == 0 {
		fmt.Println("MANY SOLUTIONS")
	} else if c < 0 {
		fmt.Println("NO SOLUTION")
	} else {
		fmt.Println((c*c - b) / a)
	}
}
