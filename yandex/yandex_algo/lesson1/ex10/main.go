package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d, e, f float32
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c, &d, &e, &f)
	op := a*d - c*b
	if op == 0 {
		if (a*d == b*c) && (a*f != b*e) {
			fmt.Print("0")
		} else {
			fmt.Print("1 ", -a/b, " ", e/b)
		}
	} else {
		fmt.Print("2 ", (e*d-f*c)/op, " ", (a*f-e*b)/op)
	}
}
