package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	side := make([]int, 3)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &side[i])
	}
	sort.Ints(side)
	if side[2] < side[1]+side[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
