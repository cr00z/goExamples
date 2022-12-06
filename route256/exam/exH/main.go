package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var results []int

func analyze(field *[][]int, lvl int) {
	maxY := len(*field)
	maxX := len((*field)[0])
	//fmt.Println(maxX, maxY)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if (*field)[y][x] == lvl*2+1 {
				results = append(results, lvl)
				topY := y
				topX := x
				botY := topY
				for ; botY < maxY && (*field)[botY][x] == lvl*2+1; botY++ {
				}
				botX := topX
				for ; botX < maxX && (*field)[y][botX] == lvl*2+1; botX++ {
				}
				area := make([][]int, botY-topY-2, botY-topY-2)
				for y2 := topY; y2 < botY; y2++ {
					if y2 > topY && y2 < botY-1 {
						area[y2-topY-1] = make([]int, botX-topX-2, botX-topX-2)
					}
					for x2 := topX; x2 < botX; x2++ {
						(*field)[y2][x2] += 2
						if y2 > topY && y2 < botY-1 && x2 > topX && x2 < botX-1 {
							area[y2-topY-1][x2-topX-1] = (*field)[y2][x2]
						}
					}
				}
				//fmt.Println(lvl, ":", topY, topX, botY, botX)
				analyze(&area, lvl+1)
			}
		}
	}
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		results = make([]int, 0)
		var y, x int
		fmt.Fscan(in, &y, &x)
		field := make([][]int, y, y)
		for j := 0; j < y; j++ {
			var line string
			fmt.Fscan(in, &line)
			field[j] = make([]int, x, x)
			for k := 0; k < x; k++ {
				if line[k] == '*' {
					field[j][k] = 1
				} else {
					field[j][k] = 0
				}
			}
		}
		analyze(&field, 0)
		sort.Ints(results)
		fmt.Println(strings.Trim(fmt.Sprint(results), "[]"))
	}
}
