package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	if n <= 6 {
		fmt.Println(1)
	} else {
		fmt.Println((n - 3) / 4 + 1)
	}
}