package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type champ struct {
	idx    int
	result int
	pos    int
}

func main() {
	var t, n, res int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(in, &n)
		champs := make([]champ, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &res)
			champs[i] = champ{
				idx:    i,
				result: res,
			}
		}
		sort.Slice(champs, func(i, j int) bool {
			return champs[i].result < champs[j].result
		})

		lastPos := 1
		champs[0].pos = lastPos

		for idx := 1; idx < len(champs); idx++ {
			if champs[idx].result-1 > champs[idx-1].result {
				lastPos = idx + 1
			}
			champs[idx].pos = lastPos
		}

		//fmt.Println(champs)

		sort.Slice(champs, func(i, j int) bool {
			return champs[i].idx < champs[j].idx
		})

		for i := 0; i < len(champs); i++ {
			fmt.Print(champs[i].pos, " ")
		}
		fmt.Println()
	}

}
