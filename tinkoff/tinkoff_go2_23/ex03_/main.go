package main

import (
	"bufio"
	"fmt"
	"os"
)

func read(in *bufio.Reader) (int, []int, []int) {
	var n int
	fmt.Fscan(in, &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	return n, s, a
}

func main() {
	in := bufio.NewReader(os.Stdin)
	n1, s1, a1 := read(in)
	_, s2, a2 := read(in)
	pos1 := 0
	pos2 := 0
	idx := 1
	result := int64(0)
	for pos1 < n1 {
		symb1 := s1[pos1]
		symb2 := s2[pos2]
		diff := a1[pos1] - a2[pos2]
		var repeat int
		if diff > 0 {
			a1[pos1] = diff
			repeat = a2[pos2]
			pos2++
		} else if diff < 0 {
			a2[pos2] = -diff
			repeat = a1[pos1]
			pos1++
		} else {
			repeat = a1[pos1]
			pos1++
			pos2++
		}
		finish := idx + repeat
		if symb1 == symb2 {
			idx = finish
		} else {
			for idx < finish {
				result += int64(idx)
				idx++
			}
		}
	}
	fmt.Println(result)
}
