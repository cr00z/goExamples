package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c, &d)
	


	if b > a {
		a, b = b, a
	}
	if d > c {
		c, d = d, c
	}
	maxAC := a
	if c > a {
		maxAC = c
	}
	maxBD := b
	if d > b {
		maxBD = d
	}
	if (a+c)*maxBD < maxAC*(b+d) {
		fmt.Println(a+c, maxBD)
	} else {
		fmt.Println(maxAC, b+d)
	}
}
