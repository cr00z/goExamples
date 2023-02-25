package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	alphabet := make(map[int32]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		for _, ch := range scanner.Text() {
			alphabet[ch]++
		}
	}

	lenAlphabet := len(alphabet)
	chars := make([]int32, 0)
	maxFreq := 0
	for key, freq := range alphabet {
		chars = append(chars, key)
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	for i := maxFreq; i > 0; i-- {
		out := strings.Builder{}
		out.Grow(lenAlphabet)
		for j := 0; j < lenAlphabet; j++ {
			if alphabet[chars[j]] < i {
				out.WriteByte(' ')
			} else {
				out.WriteByte('#')
			}
		}
		fmt.Println(out.String())
	}
	for i := 0; i < lenAlphabet; i++ {
		fmt.Print(string(chars[i]))
	}
}
