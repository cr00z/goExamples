package _11

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

func main_() {
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
		if cl[len(cl)-1].pos < n {
			cl[0].pos++
			cl[0].change = 1
		}
		var finish bool
		lenCL := len(cl)
		for i := 1; i < lenCL/2; i++ {
			if cl[i].pos-1 > cl[i-1].pos {
				cl[i].pos--
				cl[i].change = -1
				continue
			}
			if cl[i].pos == cl[i-1].pos {
				if cl[i].pos+1 < cl[i+1].pos {
					cl[i].pos++
					cl[i].change = 1
					continue
				} else {
					fmt.Println("x")
					finish = true
					break
				}
			}
			j := lenCL - i - 1
			if cl[j].pos+1 < cl[j+1].pos {
				cl[j].pos++
				cl[j].change = 1
				continue
			}
			if cl[j].pos == cl[j+1].pos {
				if cl[j].pos-1 > cl[i-1].pos {
					cl[j].pos--
					cl[j].change = -1
					continue
				} else {
					fmt.Println("x")
					finish = true
					break
				}
			}
		}
		if !finish {
			sort.Slice(cl, func(i, j int) bool {
				return cl[i].idx < cl[j].idx
			})
			for i := 0; i < lenCL; i++ {
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
