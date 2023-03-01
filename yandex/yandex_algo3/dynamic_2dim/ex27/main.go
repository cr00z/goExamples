package main

import (
	"bufio"
	"fmt"
	"os"
)

type route struct {
	cost   int
	direct byte
}

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	dp := make([][]route, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]route, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &dp[i][j].cost)
			dp[i][j].direct = 0
			max, max2 := 0, 0
			if i > 0 {
				max = dp[i-1][j].cost
			}
			if j > 0 {
				max2 = dp[i][j-1].cost
			}
			if max > max2 {
				dp[i][j].cost += max
				dp[i][j].direct = 'D'
			} else {
				dp[i][j].cost += max2
				dp[i][j].direct = 'R'
			}
		}
	}
	fmt.Println(dp[n-1][m-1].cost)
	var path string
	for n > 1 || m > 1 {
		path = string(dp[n-1][m-1].direct) + " " + path
		if dp[n-1][m-1].direct == 'R' {
			m--
		} else {
			n--
		}
	}
	fmt.Println(path)
}
