package main

import (
	"bufio"
	"fmt"
	"os"
)

func limit(num int64) (int64, int, int) {
	shift := 0
	for ;num > 9; shift++ {
		num = num / 10
	}
	result := num
	for i := 0; i < shift; i++ {
		result = result * 10 + num
	}
	return result, int(num), shift
}

func main() {
	var l, r int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &l, &r)
	
	llim, lbase, lshift := limit(l)
	if l > llim {
		lbase++
	}
	rlim, rbase, rshift := limit(r)
	if r < rlim {
		rbase--
	}
	shiftDiff := rshift - lshift
	sum := 0
	if shiftDiff == 0 {
		sum = rbase - lbase + 1
	} else if (shiftDiff == 1) {
		sum = 10 - lbase + rbase
	} else {
		sum = 10 - lbase + rbase + 9 * (shiftDiff - 1)
	}
	fmt.Println(sum)
}