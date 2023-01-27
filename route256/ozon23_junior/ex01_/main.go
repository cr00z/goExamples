package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		fmt.Println(a + b)
	}
}
