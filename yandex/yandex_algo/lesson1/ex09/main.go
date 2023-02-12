package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d, e int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c, &d, &e)
	if d > e {
		d, e = e, d
	}
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	if a <= d && b <= e {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
