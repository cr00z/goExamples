package main_

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

type proc struct {
	power  int64
	finish int64
}

func main_() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	var sumPower int64
	var n, m int
	var ai int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	procs := make([]proc, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ai)
		procs[i] = proc{ai, 0}
	}
	sort.Slice(procs, func(i, j int) bool {
		return procs[i].power < procs[j].power
	})
	minFinish := int64(0)
	current := 0
	for i := 0; i < m; i++ {
		var start, duration int64
		fmt.Fscan(in, &start, &duration)
		finish := start + duration
		if finish < minFinish {
			// если все первые процессоры заняты
			if current < n-1 {
				// берем следующий
				current++
				procs[current].finish = finish
				sumPower += procs[current].power * duration
				continue
			}
		} else {
			minFinish = int64(2_000_000_000)
			used := false
			for j := 0; j < n; j++ {
				if procs[j].finish <= start && !used {
					procs[j].finish = finish
					sumPower += procs[j].power * duration
					current = j
					used = true
				}
				if procs[j].finish < minFinish {
					minFinish = procs[j].finish
				}
			}
		}
	}
	fmt.Println(sumPower)
}
