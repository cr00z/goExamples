package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stack := make([]int64, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	for _, op := range strings.Split(line, " ") {
		if op == "" {
			break
		} else if op == "+" || op == "-" || op == "*" {
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int64
			switch op {
			case "+":
				res = op1 + op2
			case "-":
				res = op1 - op2
			case "*":
				res = op1 * op2
			}
			stack = append(stack, res)
		} else {
			num, _ := strconv.ParseInt(op, 10, 64)
			stack = append(stack, num)
		}
	}
	if len(stack) > 0 {
		fmt.Println(stack[len(stack)-1])
	}
}
