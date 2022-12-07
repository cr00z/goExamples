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
	var last int
	for i := 0; i < n; i++ {
		var item int
		fmt.Fscanln(in, &item)
		if (item > last) || (i == 0) {
			fmt.Println(item)
			last = item
		}
	}
}