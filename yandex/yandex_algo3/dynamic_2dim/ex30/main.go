package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int16, b int16) int16 {
	if b > a {
		return b
	}
	return a
}

type code struct {
	len  int16
	prev int16
}

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	seqn := make([]int16, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &seqn[i])
	}

	fmt.Fscan(in, &m)
	seqm := make([]int16, n)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &seqm[i])
	}

	dp := make([][]code, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]code, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if seqn[i-1] == seqm[j-1] {
				dp[i][j].len = dp[i-1][j-1].len + 1
				dp[i][j].prev = int16(i) - 1
			} else {
				dp[i][j].len = max(dp[i-1][j].len, dp[i][j-1].len)
			}
		}
	}

	ans := make([]int16, 0)
	i := n
	j := m
	for i > 0 && j > 0 {
		if seqn[i-1] == seqm[j-1] {
			ans = append(ans, seqn[i-1])
			i--
			j--
		} else if dp[i-1][j] == dp[i][j] {
			i--
		} else {
			j--
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Print(ans[i], " ")
	}
}
