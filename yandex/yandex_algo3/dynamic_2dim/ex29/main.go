// https://inf.1sept.ru/article.php?ID=200601603

package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	price := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &price[i])
		if price[i] > 100 {
			m++
		}
	}
	dp := make([][]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < m; j++ {
			if price[i] <= 100 {
				// покупает i-обед
				dp[i][j] = dp[i-1][j] + price[i]
				if j > 0 {
					// или тратит купон
					dp[i][j] = min(dp[i][j], dp[i-1][j-1])
				}
			} else {
				dp[i][j+1] = dp[i-1][j] + price[i]
				if j > 0 {
					// или тратит купон
					dp[i][j] = min(dp[i][j], dp[i-1][j-1])
				}
			}
		}
	}
}
