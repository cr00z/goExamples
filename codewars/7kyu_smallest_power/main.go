package main

import (
	"fmt"
	"math"
)

func FindNextPower(val, pow int) int {
	root := math.Pow(float64(val), 1/float64(pow))
	val2 := int(math.Pow(math.Ceil(root), float64(pow)))
	if val == val2 {
		val2 = int(math.Pow(math.Ceil(root)+1, float64(pow)))
	}
	return val2
}

func main() {
	fmt.Println(FindNextPower(2097152, 7))
}
