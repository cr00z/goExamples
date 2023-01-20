package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nn, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &nn)
	for i := 0; i < nn; i++ {
		fmt.Fscan(in, &n)
		if n == 3 {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
		sign := 1
		num := 1
		for j := 0; j < n; j++ {
			if n%2 == 0 {
				num = 1
			} else {
				if j%2 == 0 {
					num = (n - 3) / 2
				} else {
					num++
				}
			}
			fmt.Print(sign*num, " ")
			sign = -sign
		}
		fmt.Println()
	}
}
