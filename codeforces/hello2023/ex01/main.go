package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, ln int
	var lamps string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ln, &lamps)
		rpos := strings.IndexByte(lamps, 'R')
		lpos := strings.LastIndexByte(lamps, 'L')
		if rpos == -1 || lpos == -1 {
			fmt.Println(-1)
		} else {
			if rpos < lpos {
				fmt.Println(0)
			} else {
				fmt.Println(lpos+1)
			}
		}
	}
}