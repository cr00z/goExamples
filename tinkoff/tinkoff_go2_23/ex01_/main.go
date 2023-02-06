package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	for n > 0 {
		if m%2 == 1 && n > 0 {
			fmt.Println(m/2 + 1)
			n--
		}
		for i := m / 2; i > 0 && n > 0; i-- {
			fmt.Println(i)
			n--
			if n == 0 {
				break
			}
			fmt.Println((m + 1) - i)
			n--
		}
	}

}
