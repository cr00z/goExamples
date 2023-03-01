package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &dp[i][j])
			if i > 0 || j > 0 {
				min, min2 := 101, 101
				if i > 0 {
					min = dp[i-1][j]
				}
				if j > 0 {
					min2 = dp[i][j-1]
				}
				if min2 < min {
					min = min2
				}
				dp[i][j] += min
			}
		}
	}
	fmt.Println(dp[n-1][m-1])
}
