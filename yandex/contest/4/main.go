package main

import (
	"bufio"
	"fmt"
	"os"
)

func braces(count int, str string, left int, right int) {
	if left == count && right == count {
		fmt.Println(str)
	} else {
		if left < count {
			braces(count, str + "(", left + 1, right)
		}
		if right < left {
			braces(count, str + ")", left, right + 1)
		}
	}
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &n)
	braces(n, "", 0, 0)
}