package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Letter struct {
	ch rune
	num int
}

func main() {
	var a string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a)

	chs := make(map[rune]int, len(a))
	maxlen := 0
	for _, ch := range a {
		chs[ch] += 1
		if chs[ch] > maxlen {
			maxlen = chs[ch]
		}
	}

	if len(a) < 2 * maxlen {
		fmt.Println("IMPOSSIBLE")
		return
	}
	
	lts := make([]Letter, len(chs))
	for key, val := range chs {
		lts = append(lts, Letter{key, val})
	}
	sort.Slice(lts, func(i int, j int)bool {
		return lts[i].num > lts[j].num
	})

	for _, lt := range lts {
		fmt.Print(strings.Repeat(string(lt.ch), lt.num))
	}
	fmt.Println()
	
	for i, lt := range lts {
		if i != 0 {
			fmt.Print(strings.Repeat(string(lt.ch), lt.num))
		}
	}
	fmt.Print(strings.Repeat(string(lts[0].ch), lts[0].num))
}