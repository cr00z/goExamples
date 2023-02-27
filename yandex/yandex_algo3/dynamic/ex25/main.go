package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	emp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &emp[i])
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		out.WriteString(strconv.Itoa(emp[i]))
		out.WriteByte(10)
	}
}
