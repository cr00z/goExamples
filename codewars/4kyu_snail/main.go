package main

import (
	"fmt"
)

func Snail(snailMap [][]int) []int {
	size := len(snailMap)
	result := []int{}
	if size == 0 || ((size == 1) && (len(snailMap[0]) == 0)) {
		return result
	}
	stop := size/2 + int(size%2)
	for i := 0; i < stop; i++ {
		for j := i; j < size-i; j++ {
			result = append(result, snailMap[i][j])
		}
		for j := i + 1; j < size-i; j++ {
			result = append(result, snailMap[j][size-i-1])
		}
		for j := i + 1; j < size-i; j++ {
			result = append(result, snailMap[size-i-1][size-j-1])
		}
		for j := i + 1; j < size-i-1; j++ {
			result = append(result, snailMap[size-j-1][i])
		}
	}
	return result
}

func main() {
	fmt.Println(Snail([][]int{}))
}
