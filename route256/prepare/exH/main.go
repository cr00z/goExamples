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
	min_stopped := 2000000000
	p := 0
	for i := 0; i < workNum; i++ {
		var start, duration int
		fmt.Fscan(in, &start, &duration)
		//fmt.Println("Work:", start, duration)
		for cnt := 0; cnt < procNum; cnt++ {
			if p == procNum {
				p = 0
			}
			if max_stopped < start {

			}
			if stopped[p] < start {
				energy += int64(duration) * int64(procs[p])
				stopped[p] = start + duration - 1
				if stopped[p] > max_stopped {
					min_stopped = stopped[p]
				}
				break
			}
			p++
		}
		//fmt.Println(procs)
		//fmt.Println(stopped)
	}
	fmt.Println(energy)
}
