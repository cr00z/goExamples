package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var wordsNum, reqNum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &wordsNum)
	dict := make([]string, wordsNum)
	for i := 0; i < wordsNum; i++ {
		var word string
		fmt.Fscan(in, &word)
		dict[i] = Reverse(word)
	}
	sort.Strings(dict)
	//fmt.Println(dict)
	fmt.Fscan(in, &reqNum)
	for i := 0; i < reqNum; i++ {
		var request string
		fmt.Fscan(in, &request)
		request = Reverse(request)
		//suffix := string(request[0])
		var cand, finded int
		for cand < len(dict) && dict[cand][0] < request[0] {
			finded = cand
			cand++
		}
		pos := 1
		if dict[0] == request {
			cand++
			finded++
		}
		for {
			if cand == len(dict) {
				break
			}
			if pos > len(request) {
				break
			}
			prefix := request[0:pos]
			if dict[cand] != request && strings.HasPrefix(dict[cand], prefix) {
				finded = cand
				pos++
			} else {
				cand++
			}
		}
		fmt.Println(Reverse(dict[finded]))
	}
}
