package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Record struct {
	time int
	id int
	status string
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	items := make([]Record, n, n)
	for i := 0; i < n; i++ {
		var d, h, m int
		fmt.Fscan(in, &d, &h, &m)
		items[i].time = d * 24 * 60 + h * 60 + m
		fmt.Fscan(in, &items[i].id)
		fmt.Fscan(in, &items[i].status)
	}
	sort.Slice(items, func(i, j int)bool {
		return items[i].time < items[j].time
	})

	rockets := make(map[int]int)
	rocketsSum := make(map[int]int)
	for i := 0; i < n; i++ {
		switch items[i].status {
		case "A":
			rockets[items[i].id] = items[i].time
		// case "B":
		// 	rockets[items[i].id] = items[i].time - rockets[items[i].id]
		case "S":
			fallthrough
		case "C":
			rocketsSum[items[i].id] += items[i].time - rockets[items[i].id]
		}
	}


	keys := make([]int, 0, len(rocketsSum))
	for key := range rocketsSum {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(rocketsSum[key])
	}
}