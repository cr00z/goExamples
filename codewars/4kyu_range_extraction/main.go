package main

import (
	"fmt"
	"strconv"
)

func Solution(list []int) string {
	var result string
	for i := 0; i < len(list); i++ {
		if i == len(list)-1 {
			result += strconv.Itoa(list[i])
		}
		for j := i + 1; j < len(list); j++ {
			if list[i] != list[j]-j+i {
				if j == i+1 || j == i+2 {
					result += strconv.Itoa(list[i]) + ","
				} else {
					result += strconv.Itoa(list[i]) + "-" + strconv.Itoa(list[j-1]) + ","
					i = j - 1
				}
				break
			} else if j == len(list)-1 {
				if j == i+1 {
					result += strconv.Itoa(list[i]) + "," + strconv.Itoa(list[j])
				} else {
					result += strconv.Itoa(list[i]) + "-" + strconv.Itoa(list[j])
				}
				i = j
			}
		}
	}
	return result
}

func main() {
	fmt.Println(Solution([]int{-6, -5, -4}))
}
