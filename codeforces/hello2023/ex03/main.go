package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	var nn, n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &nn)
	for i := 0; i < nn; i++ {
		fmt.Fscan(in, &n, &m)
		for i := 0; i < n; i++ {


		emp := make([]int, n, n)
		fmt.Fscan(in, &emp[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println(emp[i])
	}
}