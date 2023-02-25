package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stack := make([]int, 0)

	var cmd string
	in := bufio.NewReader(os.Stdin)
LOOP:
	for {
		fmt.Fscan(in, &cmd)
		switch cmd {
		case "push":
			var n int
			fmt.Fscan(in, &n)
			stack = append(stack, n)
			fmt.Println("ok")

		case "pop":
			if len(stack) > 0 {
				last := stack[0]
				stack = stack[1:len(stack)]
				fmt.Println(last)
			} else {
				fmt.Println("error")
			}

		case "front":
			if len(stack) > 0 {
				fmt.Println(stack[0])
			} else {
				fmt.Println("error")
			}

		case "size":
			fmt.Println(len(stack))

		case "clear":
			stack = make([]int, 0)
			fmt.Println("ok")

		case "exit":
			fmt.Println("bye")
			break LOOP
		}
	}
}
