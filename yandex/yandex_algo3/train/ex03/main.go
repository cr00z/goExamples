package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getCount(sticks []int, find int) int {
	l := len(sticks)

	if find <= sticks[0] {
		return 0
	}
	if find > sticks[l-1] {
		return l
	}

	pos := l / 2
	oldPos := 0
	oldOldPos := 0
	for oldPos != pos && oldOldPos != pos {
		oldOldPos = oldPos
		oldPos = pos
		if sticks[pos] < find {
			pos = (pos + l) / 2
		} else {
			pos = pos / 2
		}
	}

	return pos + 1
}

func main() {
	var n, k, stick, find int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	sticksUniq := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &stick)
		sticksUniq[stick] = struct{}{}
	}

	sticks := make([]int, len(sticksUniq))
	idx := 0
	for st := range sticksUniq {
		sticks[idx] = st
		idx++
	}
	sort.Ints(sticks)

	fmt.Fscan(in, &k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &find)
		fmt.Println(getCount(sticks, find))
	}
}
