package main

import (
	"bufio"
	"fmt"
	"os"
)

func formCommand(devs []int) [][2]int {
	command := make([][2]int, 0)
	for dev1Idx := 0; dev1Idx < len(devs)-1; dev1Idx++ {
		if devs[dev1Idx] == 0 {
			continue
		}
		lvlDiff := 100
		var candidate int
		for dev2Idx := dev1Idx + 1; dev2Idx < len(devs); dev2Idx++ {
			if devs[dev2Idx] == 0 {
				continue
			}
			currDiff := devs[dev1Idx] - devs[dev2Idx]
			if currDiff < 0 {
				currDiff = -currDiff
			}
			if currDiff < lvlDiff {
				candidate = dev2Idx
				lvlDiff = currDiff
			}
		}
		command = append(command, [2]int{dev1Idx + 1, candidate + 1})
		devs[candidate] = 0
	}
	return command
}

func main() {
	var cmdsNum, devNum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &cmdsNum)
	for i := 0; i < cmdsNum; i++ {
		fmt.Fscan(in, &devNum)
		devs := make([]int, devNum)
		for j := 0; j < devNum; j++ {
			fmt.Fscan(in, &devs[j])
		}
		command := formCommand(devs)
		for j := 0; j < devNum/2; j++ {
			fmt.Println(command[j][0], command[j][1])
		}
		if i < cmdsNum-1 {
			fmt.Println()
		}
	}
}
