package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	var str string
	for k := 0; k < n; k++ {
		fmt.Fscan(in, &str)
		items := strings.Split(str, ",")
		chs := make(map[rune]struct{})
		for _, ch := range items[0]+items[1]+items[2] {
			chs[ch] = struct{}{}
		}
		var sum int
		for _, ch := range items[3]+items[4] {
			sum += int(ch) - '0'
		}
		sum *= 64
		first := int(items[0][0] - 'A' + 1) * 256
		s1 := len(chs) + sum + first
		h := fmt.Sprintf("%X", s1)
		for len(h) < 3 {
			h = "0" + h
		}
		fmt.Println(h[len(h)-3:])
	}
}