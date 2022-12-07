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
	var i int
	for i = n / 2; (i > 1) && (n % i != 0); i-- {}
	fmt.Println(i, n - i)
}