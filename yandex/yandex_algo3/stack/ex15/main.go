package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	towns := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &towns[i])
	}

	pos := make([]int, n)
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if towns[i] > towns[j] {
				pos[i] = j
				break
			}
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		out.WriteString(strconv.Itoa(pos[i]))
		out.WriteByte(32)
	}
}
