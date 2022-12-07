package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &a, &b, &c, &d)
	res := a
	if d > b {
		res += (d - b) * c
	}
	fmt.Println(res)
}