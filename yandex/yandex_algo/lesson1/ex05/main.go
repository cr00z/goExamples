package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k1, m, k2, p2, n2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &k1, &m, &k2, &p2, &n2)
	fullN := (p2-1)*m + n2
	kk := k2 / fullN
	if k2 != kk*fullN {
		kk++
	}
	p1 := k1/(m*kk) + 1
	fmt.Println(p1)
	n1 := (k1 - (p1-1)*m) / kk
	if k1-(p1-1)*m != n1*kk {
		n1++
	}
	fmt.Println(n1)
}
