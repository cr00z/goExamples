package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	ai := make([]int, n + 1, n + 1)
	ai[0] = -2_000_000_000
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &ai[i])
	}
	sort.Ints(ai)
	ai = append(ai, -2_000_000_000)

	if n == 1 {
		fmt.Println("NO")
		return
	}

	var closed bool
LOOP:
	for i := 1; i <= n && !closed; i++ {
		if ai[i] + ai[i+1] != k {
			continue
		}
		for j := 1; j <= n; j++ {
			if j == i || j == i+1 {
				continue
			}
			if (ai[j] + ai[i+1] != k) && (ai[i] + ai[j-1] != k) && (ai[i] + ai[j+1] != k) {
				ai[i], ai[j] = ai[j], ai[i]
				continue LOOP
			}
		}
	}

	if closed {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(strings.Trim(fmt.Sprint(ai[1:len(ai)-1]), "[]"))
	}
}