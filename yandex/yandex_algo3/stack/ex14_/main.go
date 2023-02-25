package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	deadEnd := make([]int, 0)
	current := 1

	var n, carr int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &carr)

		for len(deadEnd) > 0 {
			if deadEnd[len(deadEnd)-1] == current {
				deadEnd = deadEnd[:len(deadEnd)-1]
				current++
			} else {
				break
			}
		}

		if carr == current {
			current++
		} else {
			deadEnd = append(deadEnd, carr)
		}
	}

	for len(deadEnd) > 0 {
		if deadEnd[len(deadEnd)-1] == current {
			deadEnd = deadEnd[:len(deadEnd)-1]
			current++
		} else {
			break
		}
	}

	if len(deadEnd) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
