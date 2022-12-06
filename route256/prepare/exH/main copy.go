package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var procNum, workNum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &procNum, &workNum)
	procs := make([]int, procNum)
	stopped := make([]int, procNum)
	for i := 0; i < procNum; i++ {
		fmt.Fscan(in, &procs[i])
	}
	sort.Ints(procs)
	var energy int64
	for i := 0; i < workNum; i++ {
		var start, duration int
		fmt.Fscan(in, &start, &duration)
		//fmt.Println("Work:", start, duration)
		for p := 0; p < procNum; p++ {
			if stopped[p] < start {
				energy += int64(duration) * int64(procs[p])
				stopped[p] = start + duration - 1
				break
			}
		}
		//fmt.Println(procs)
		//fmt.Println(stopped)
	}
	fmt.Println(energy)
}
