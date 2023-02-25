package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, num, min int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k, &min)
	for i := 1; i < k; i++ {
		fmt.Fscan(in, &num)
		if num < min {
			min = num
		}
	}
	fmt.Println(min)
	for i := 0; i < n-k; i++ {
		fmt.Fscan(in, &num) {

		}
	}
}