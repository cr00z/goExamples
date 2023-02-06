package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, ship int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		ships := make(map[int]int, n)
		for j := 0; j < 10; j++ {
			fmt.Fscan(in, &ship)
			ships[ship]++
		}
		if ships[1] == 4 && ships[2] == 3 && ships[3] == 2 && ships[4] == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}