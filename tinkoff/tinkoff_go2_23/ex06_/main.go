package main

import (
	"bufio"
	"fmt"
	"os"
)

const stubValue = 10_000_000

func main() {
	var n, x, ai int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &x)

	dp := make([]int, x+1)
	dp[0] = 0
	for j := 1; j <= x; j++ {
		dp[j] = stubValue
	}

	// массив для накопительной суммы кристаллов
	a := make([]int, n+1)
	sum, i := 0, 1
	for ; i <= n; i++ {
		fmt.Fscan(in, &ai)
		sum += ai
		a[i] = sum
		// сразу забиваем массив динамического прораммирования значениями
		if sum <= len(dp) {
			dp[sum] = i
		}
	}
	// добиваем до конца массива ai-ми
	sum += ai
	for sum <= len(dp) {
		dp[sum] = i
		i++
		sum += ai
	}

	// пересчитываем массив
	for j := 1; j <= x; j++ {
	LOOP:
		for i := 1; i <= n; i++ {
			if a[i] > j {
				break LOOP
			}
			if dp[j] > dp[j-a[i]]+i+1 {
				dp[j] = dp[j-a[i]] + i + 1
			}
		}
		// дошли до конца массива дней - досчитываем допами
		if i == n {
			dop := 1
			for j-a[i]-ai*dop >= 0 {
				if dp[j] > dp[j-a[i]-ai*dop]+i+1+dop {
					dp[j] = dp[j-a[i]-ai*dop] + i + 1 + dop
				}
				dop++
			}
		}
	}

	if dp[x] == stubValue {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[x])
	}
}
