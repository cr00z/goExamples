package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &a)
	var i int64 = 1
	cut := 0
	for {
		if i > a - 1 {
			break
		}
		i *= 2
		cut += 1
	}
	fmt.Println(cut)
}