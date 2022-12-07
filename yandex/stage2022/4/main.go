package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Order struct {
	start int
	end int
	cost int
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	itemsSt := make([]Order, n, n)
	itemsEnd := make([]Order, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &itemsSt[i].start, &itemsSt[i].end, &itemsSt[i].cost)
		itemsEnd[i] = itemsSt[i]
	}
	sort.Slice(itemsSt, func(i, j int)bool {
		return itemsSt[i].start < itemsSt[j].start
	})
	sort.Slice(itemsEnd, func(i, j int)bool {
		return itemsEnd[i].end < itemsEnd[j].end
	})

	var m, start, stop, tp int
	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &start, &stop, &tp)
		//fmt.Println(start, stop, tp)
		var cost int64
		if tp == 1 {

			startIdx := 0
			stopIdx := n
			idx := (stopIdx - startIdx) / 2
			for itemsSt[idx].start >= start && idx != 0 {
				stopIdx = idx
				idx = (stopIdx - startIdx) / 2
			}

			for i := idx; i < n; i++ {
				if itemsSt[i].start < start {
					continue
				}
				if itemsSt[i].start >= start && itemsSt[i].start <= stop {
					cost += int64(itemsSt[i].cost)
				}
				if itemsSt[i].start > stop {
					break
				}
			}
			fmt.Println(cost)
		} else {

			startIdx := 0
			stopIdx := n
			idx := (stopIdx - startIdx) / 2
			for itemsEnd[idx].end >= start && idx != 0 {
				stopIdx = idx
				idx = (stopIdx - startIdx) / 2
			}		

			for i := idx; i < n; i++ {
				if itemsEnd[i].end < start {
					continue
				}
				if itemsEnd[i].end >= start && itemsEnd[i].end <= stop {
					cost += int64(itemsEnd[i].end - itemsEnd[i].start)
				}
				if itemsEnd[i].end > stop {
					break
				}
			}
			fmt.Println(cost)
		}
	}


	// fmt.Println(n)
	// for i := 0; i < n; i++ {
	// 	fmt.Println(itemsSt[i])
	// }
	// for i := 0; i < n; i++ {
	// 	fmt.Println(itemsEnd[i])
	// }
}