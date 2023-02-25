package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeSum(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	result := make([][]int, n/2)
	for i := 0; i < n/2; i++ {
		result[i] = make([]int, m/2)
		for j := 0; j < m/2; j++ {
			result[i][j] = matrix[i*2][j*2] + matrix[i*2][j*2+1] +
				matrix[i*2+1][j*2] + matrix[i*2+1][j*2+1]
		}
	}
	return result
}

func main() {
	var n, m, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m, &k)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}
	matrix2 := makeSum(matrix)
	matrix4 := makeSum(matrix2)
	fmt.Println(matrix4)
}
