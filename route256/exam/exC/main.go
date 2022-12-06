package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var code string
		fmt.Fscan(in, &code)
		pos := 0
		for {
			if code[pos] == '0' {
				pos += 2
				fmt.Print("a")
			} else {
				pos++
				if code[pos] == '0' {
					pos++
					if code[pos] == '0' {
						fmt.Print("b")
					} else {
						fmt.Print("c")
					}
				} else {
					fmt.Print("d")
				}
				pos++
			}
			if pos == len(code) {
				break
			}
		}
		fmt.Println()
	}
}
