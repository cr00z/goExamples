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
	emp := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &emp[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println(emp[i])
	}
}