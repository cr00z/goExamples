package main

import (
	"bufio"
	"fmt"
	"os"
)

// func isEffective(report *[]int) bool {
// 	reportLen := len(*report)
// 	for i := 0; i < reportLen; i++ {
// 		j := i + 1
// 		for j < reportLen && (*report)[i] == (*report)[j] {
// 			j++
// 		}
// 		i = j - 1
// 		for k := j; k < reportLen; k++ {
// 			if (*report)[i] == (*report)[k] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func isEffective(report *[]int) bool {
	checked := make(map[int]bool)
	reportLen := len(*report)
	for i := 0; i < reportLen; i++ {
		j := i + 1
		for j < reportLen && (*report)[i] == (*report)[j] {
			j++
		}
		i = j - 1
		if _, inMap := checked[(*report)[i]]; inMap {
			return false
		} else {
			checked[(*report)[i]] = true
		}
	}
	return true
}

func main() {
	var workersNum, worksNum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &workersNum)
	for i := 0; i < workersNum; i++ {
		fmt.Fscan(in, &worksNum)
		report := make([]int, worksNum)
		for j := 0; j < worksNum; j++ {
			fmt.Fscan(in, &report[j])
		}
		if isEffective(&report) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
