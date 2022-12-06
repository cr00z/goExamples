package main

import (
	"fmt"
	"math"
	"sort"
)

func makeNum(digits []int) int {
	result := 0
	for pos, digit := range digits {
		result += digit * int(math.Pow10(pos))
	}
	return result
}

func getDigits(n int) []int {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	return digits
}

func NextBigger(n int) int {
	digits := getDigits(n)
	success := false
	for i := 0; i < len(digits) && !success; i++ {
		for j := i + 1; j < len(digits) && !success; j++ {
			if digits[i] > digits[j] {
				candidate1 := makeNum(digits[:j])
				candidate2 := NextBigger(candidate1)
				if candidate2 > candidate1 {
					digits = append(getDigits(candidate2), digits[j:]...)
				} else {
					digits[i], digits[j] = digits[j], digits[i]
					sort.Slice(digits[0:j], func(x, y int) bool {
						return digits[0:j][x] > digits[0:j][y]
					})
				}
				success = true
			}
		}
	}
	result := -1
	if success {
		result = makeNum(digits)
	}
	return result
}

func main() {
	fmt.Println(NextBigger(59884848459853))
}
