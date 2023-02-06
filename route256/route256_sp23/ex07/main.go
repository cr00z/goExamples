package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pat struct {
	idx    int
	pos    int
	change int
}

func main() {
	var t, n, m, pos int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(in, &n, &m)
		cl := make([]*pat, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &pos)
			cl[i] = &pat{
				idx: i,
				pos: pos,
			}
		}
		sort.Slice(cl, func(i, j int) bool {
			return cl[i].pos < cl[j].pos
		})
		if cl[0].pos > 1 {
			cl[0].pos--
			cl[0].change = -1
		}
		for i := 1; i < m; i++ {
			if cl[i].pos-1 > cl[i-1].pos {
				cl[i].pos--
				cl[i].change = -1
			}
		}
		if cl[m-1].pos < n && cl[m-1].change == -1 {
			cl[m-1].pos++
			cl[m-1].change = 0
		}
		if cl[m-1].pos < n && cl[m-1].change == 0 {
			cl[m-1].pos++
			cl[m-1].change = 1
		}
		for i := m - 2; i > 0; i-- {
			if cl[i].pos+1 < cl[i+1].pos && cl[i].change == -1 {
				cl[i].pos++
				cl[i].change = 0
			}
			if cl[i].pos+1 < cl[i+1].pos && cl[i].change == 0 {
				cl[i].pos++
				cl[i].change = 1
			}
		}
		var failure bool
		for i := 1; i < m; i++ {
			if cl[i].pos == cl[i-1].pos {
				fmt.Println("x")
				failure = true
				break
			}
		}
		if !failure {
			sort.Slice(cl, func(i, j int) bool {
				return cl[i].idx < cl[j].idx
			})
			for i := 0; i < m; i++ {
				if cl[i].change == -1 {
					fmt.Print("-")
				} else if cl[i].change == 1 {
					fmt.Print("+")
				} else {
					fmt.Print("0")
				}
			}
			fmt.Println()
		}
	}
}
