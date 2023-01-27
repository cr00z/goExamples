package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t, n, m, k, ck int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for it := 0; it < t; it++ {
		fmt.Fscan(in, &n, &m)
		table := make([]*[]int, n, n)
		for i := 0; i < n; i++ {
			row := make([]int, m, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &row[j])
			}
			table[i] = &row
		}
		fmt.Fscan(in, &k)
		for ki := 0; ki < k; ki++ {
			fmt.Fscan(in, &ck)
			ck--
			sort.SliceStable(table, func(i, j int) bool {
				return (*table[i])[ck] < (*table[j])[ck]
			})
		}
		// print
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Print((*table[i])[j])
				if j != m-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
