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