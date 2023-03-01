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
	}
	if n > 1 && m > 2 {
		dp[1][2] = 1
	}
	if n > 2 && m > 1 {
		dp[2][1] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i-1 > 0 && j > 0 {
				dp[i][j] += dp[i-2][j-1]
			}
			if i > 0 && j-1 > 0 {
				dp[i][j] += dp[i-1][j-2]
			}
		}
	}
	fmt.Println(dp[n-1][m-1])
}
