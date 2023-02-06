package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	var str string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(in, &n)
		vocab := make(map[string]struct{}, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &str)
			last := byte(0)
			repeat := 1
			result := make([]byte, 0, len(str))
			for idx := range str {
				if str[idx] != last {
					repeat = 1
					result = append(result, str[idx])
				} else {
					if repeat == 1 {
						result = append(result, str[idx])
					}
					repeat++
				}
				last = str[idx]
			}
			vocab[string(result)] = struct{}{}
		}
		fmt.Println(len(vocab))
	}
}
