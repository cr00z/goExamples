package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	lines := make([][]byte, 0, n)
	var line string
	fmt.Fscanln(in, &line)
	for i := 0; i < n; i++ {
		fmt.Fscanln(in, &line)
		lines = append(lines, []byte(line))
	}

	fmt.Fscan(in, &m)
	fmt.Fscanln(in, &line)
	var np, startpos int
	var side, pos string
	for i := 0; i < m; i++ {
		fmt.Fscanf(in, "%d %s %s\n", &np, &side, &pos)
		//fmt.Println(np, side, pos)
		if side == "left" {
			if pos == "window" {
				startpos = 0
			} else {
				startpos = 3 - np
			}
		} else {
			if pos == "window" {
				startpos = 7 - np
			} else {
				startpos = 4
			}
		}
		
		var seat bool
		seats := strings.Repeat(".", np)
		i := 0
		for ; i < n; i++ {
			if string(lines[i][startpos:startpos+np]) == seats {
				seat = true
				for j := startpos; j < startpos+np; j++ {
					lines[i][j] = 'X'
				}
				break
			}
		}

		if seat {
			fmt.Print("Passengers can take seats:")
			for j := 0; j < np; j++ {
				shift := 0
				if startpos > 3 {
					shift = 1
				}
				fmt.Printf(" %d%s", i+1, string('A'+startpos+j-shift))
			}
			fmt.Println()
			for k := 0; k < n; k++ {
				fmt.Println(string(lines[k]))
			}
			for j := startpos; j < startpos+np; j++ {
				lines[i][j] = '#'
			}
		} else {
			fmt.Println("Cannot fulfill passengers requirements")
		}
	}
}