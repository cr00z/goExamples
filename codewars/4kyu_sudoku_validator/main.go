package main

import "fmt"

func ValidateSolution(m [][]int) bool {
	result := true
	for i := 0; (i < 9) && result; i++ {
		test := 0
		for j := 0; j < 9; j++ {
			test += 2 << m[i][j]
			test += 2 << (m[j][i] + 9)
		}
		if test != 1048572 {
			result = false
		}
	}
	for x := 0; (x < 3) && result; x++ {
		for y := 0; (y < 3) && result; y++ {
			test := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					test += 2 << m[x*3+i][y*3+j]
				}
			}
			if test != 2044 {
				result = false
			}
		}
	}
	return result
}

func main() {
	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	var testFalse = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{3, 1, 2, 6, 4, 5, 9, 7, 8},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 4, 8, 9, 7, 2, 3, 1},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 7, 2, 3, 1, 5, 6, 4},
		{9, 7, 8, 3, 1, 2, 6, 4, 5},
	}
	fmt.Println(ValidateSolution(testTrue))
	fmt.Println(ValidateSolution(testFalse))
}
