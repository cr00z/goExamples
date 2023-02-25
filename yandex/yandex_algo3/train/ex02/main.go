package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countBeauty(line string, num int) int {
	maxCount := 0
	lenLine := len(line)

	for i := 0; i < lenLine-maxCount; i++ {
		dupl := 0
		count := 0
	LOOP:
		for j := i + 1; j < lenLine; j++ {
			if line[i] == line[j] {
				dupl++
			} else if count < num {
				count++
			}
			if count == num {
				j++
				for j < lenLine {
					if line[i] == line[j] {
						dupl++
					}
					j++
				}
				break LOOP
			}
		}

		if count+dupl+1 > maxCount {
			maxCount = count + dupl + 1
		}
	}

	return maxCount
}

func main() {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	num, _ := strconv.Atoi(myscanner.Text())
	myscanner.Scan()
	line := myscanner.Text()

	if len(line) == 0 {
		fmt.Println(0)
		os.Exit(0)
	}

	lenLine := len(line)
	if num > lenLine-1 {
		num = lenLine - 1
	}

	maxCount := countBeauty(line, num)
	maxCount2 := countBeauty(Reverse(line), num)
	if maxCount2 > maxCount {
		maxCount = maxCount2
	}

	fmt.Println(maxCount)
}
