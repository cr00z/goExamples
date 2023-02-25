package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	line := myscanner.Text()

	letters := make(map[rune]int64)
	lenLine := len(line)
	for idx, ltr := range line {
		letters[ltr] += int64(idx+1) * int64(lenLine-idx)
	}

	keys := make([]rune, 0, len(letters))
	for key := range letters {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, ch := range keys {
		fmt.Print(string(ch), ": ", letters[ch], "\n")
	}
}
