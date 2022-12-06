package main

import (
	"bufio"
	"fmt"
	"os"
)

func fillMap(hexMap *[][]byte, pos [2]int, idx byte) bool {
	(*hexMap)[pos[0]][pos[1]] = 0
	lenY := len(*hexMap)
	lenX := len((*hexMap)[0])
	newPos := [6][2]int{
		{0, -1},
		{-1, -1 + (pos[0] % 2)},
		{-1, (pos[0] % 2)},
		{0, 1},
		{1, -1 + (pos[0] % 2)},
		{1, (pos[0] % 2)},
	}
	//fmt.Println(newPos)
	result := true
	for _, nPos := range newPos {
		if pos[0]+nPos[0] < 0 || pos[0]+nPos[0] == lenY || pos[1]+nPos[1] < 0 || pos[1]+nPos[1] == lenX {
			continue
		}
		if (*hexMap)[pos[0]+nPos[0]][pos[1]+nPos[1]] == 0 {
			continue
		}
		if (*hexMap)[pos[0]+nPos[0]][pos[1]+nPos[1]] > idx {
			continue
		}
		if (*hexMap)[pos[0]+nPos[0]][pos[1]+nPos[1]] < idx {
			return false
		}
		result = result && fillMap(hexMap, [2]int{pos[0] + nPos[0], pos[1] + nPos[1]}, idx)
	}
	return result
}

func checkMap(hexMap *[][]byte, startPos [27][2]int) bool {
	for idx, pos := range startPos {
		if pos[0] == -1 {
			continue
		}
		//fmt.Println(string(idx+'@'), pos)
		if !fillMap(hexMap, pos, byte(idx)) {
			return false
		}
	}
	sum := 0
	for y := range *hexMap {
		for x := range (*hexMap)[y] {
			sum += int((*hexMap)[y][x])
		}
	}
	if sum > 0 {
		return false
	}
	return true
}

func printMap(hexMap *[][]byte) {
	for idx, str := range *hexMap {
		if idx%2 == 1 {
			fmt.Print(".")
		}
		for _, char := range str {
			fmt.Print(string(char+'@'), ".")
		}
		fmt.Println()
	}
}

func main() {
	var mapsNum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &mapsNum)
	for mapNum := 0; mapNum < mapsNum; mapNum++ {
		var numStr, numSymb int
		fmt.Fscan(in, &numStr, &numSymb)
		hexMap := make([][]byte, numStr)
		startPos := [27][2]int{}
		for idx := 0; idx < 27; idx++ {
			startPos[idx] = [2]int{-1, -1}
		}
		var mapStr string
		for idx := 0; idx < numStr; idx++ {
			hexMap[idx] = make([]byte, numSymb/2+(numSymb%2))
			fmt.Fscan(in, &mapStr)
			for i := idx % 2; i < numSymb; i += 2 {
				char := mapStr[i] - '@'
				hexMap[idx][i/2] = char
				startPos[char] = [2]int{idx, i / 2}
			}
		}
		//printMap(&hexMap)
		if checkMap(&hexMap, startPos) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		//printMap(&hexMap)
	}
}
