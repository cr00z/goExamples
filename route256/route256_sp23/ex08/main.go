package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var str string
	in := bufio.NewReader(os.Stdin)
	field := make([][]byte, 8)
	for i := 0; i < 8; i++ {
		fmt.Fscan(in, &str)
		field[i] = []byte(str)
	}
	var n, l, lenF int
	fmt.Fscan(in, &n, &l)
	fig := make([][]byte, l)
	for i := 0; i < l; i++ {
		fmt.Fscan(in, &str)

		}
	}
	emp := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &emp[i])
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < n; i++ {
		writer.WriteString(strconv.Itoa(emp[i])) // запись строки
		writer.WriteString("\n")                 // перевод строки
	}

}
