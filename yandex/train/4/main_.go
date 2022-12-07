package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_() {
	var divisor int64 = 1_000_000_007
	var n int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	st := n - 2
	var mul int64 = 1
	for st > 1 {

		if st % 2 == 1 {
			mul = mul * n
			st--
		}
		n = n * n
		st = st / 2
		if n > divisor {
			n = n % divisor
		}
		if mul > divisor {
			mul = mul % divisor
		}
		//fmt.Println("St:", st, "N:", n, "Mul:", mul)
	}
	mul = mul * n
	if mul > divisor {
		mul = mul % divisor
	}
	fmt.Println(mul)
}