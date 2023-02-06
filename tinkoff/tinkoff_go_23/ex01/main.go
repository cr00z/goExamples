package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var str string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &str)
LOOP:
	for i := 0; i < n; i++ {
		fmt.Fscanln(in, &str)
		symbs := make(map[byte]int)
		for idx := range str {
			symbs[str[idx]]++
		}
		if len(symbs) != 2 {
			fmt.Println("No")
			continue
		}
		for symb := range symbs {
			if symbs[symb] != 2 {
				fmt.Println("No")
				continue LOOP
			}
		}
		fmt.Println("Yes")
	}
}
