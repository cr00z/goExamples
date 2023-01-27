package main

import (
	"bufio"
	"fmt"
	"os"
)

func fillField(field [][]byte, color byte, i int, j int) {
	if field[i][j] != color {
		return
	}
	maxI := len(field) - 1
	maxJ := len(field[i]) - 1
	field[i][j] = 0
	if i%2 == 1 {
		if i > 0 {
			fillField(field, color, i-1, j)
			if j < maxJ {
				fillField(field, color, i-1, j+1)
			}
		}
		if j > 0 {
			fillField(field, color, i, j-1)
		}
		if j < maxJ {
			fillField(field, color, i, j+1)
		}
		if i < maxI {
			fillField(field, color, i+1, j)
			if j < maxJ {
				fillField(field, color, i+1, j+1)
			}
		}
	} else {
		if i > 0 {
			if j > 0 {
				fillField(field, color, i-1, j-1)
			}
			fillField(field, color, i-1, j)
		}
		if j > 0 {
			fillField(field, color, i, j-1)
		}
		if j < maxJ {
			fillField(field, color, i, j+1)
		}
		if i < maxI {
			if j > 0 {
				fillField(field, color, i+1, j-1)
			}
			fillField(field, color, i+1, j)
		}
	}
}

func main() {
	var t, n, m int
	var str string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	in.ReadString('\n')
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(in, &n, &m)
		in.ReadString('\n')

		field := make([][]byte, n, n)
		for i := 0; i < n; i++ {
			str, _ = in.ReadString('\n')
			field[i] = make([]byte, m/2+1, m/2+1)
			for j := i % 2; j < m; j += 2 {
				field[i][j/2] = str[j]
			}
		}

		result := "YES"
		colors := make(map[byte]struct{})
	LOOP:
		for i := 0; i < n; i++ {
			for j := 0; j < m/2+1; j++ {
				color := field[i][j]
				if color == 0 {
					continue
				}
				if _, inMap := colors[color]; inMap {
					result = "NO"
					break LOOP
				}
				colors[color] = struct{}{}
				fillField(field, color, i, j)
			}
		}

		//for i := 0; i < n; i++ {
		//	fmt.Println(field[i])
		//}
		fmt.Println(result)
	}
}
