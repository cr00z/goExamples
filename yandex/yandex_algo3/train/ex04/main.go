package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, lineP, lrP int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k, &lineP, &lrP)
	posP := (lineP-1)*2 + lrP // место пети от 1 до n
	//fullN := n + (n % 2)			// максимальное место в классе
	posV := posP + k // позиция васи сзади
	if posV <= n {
		fmt.Println((posV-1)/2+1, (posV-1)%2+1)
	} else {
		posV = posP - k
		if posV > 0 {
			fmt.Println((posV-1)/2+1, (posV-1)%2+1)
		} else {
			fmt.Println(-1)
		}
	}
}
