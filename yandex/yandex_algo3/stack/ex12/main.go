package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stack := make([]rune, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	var correct = true

	for _, ch := range line {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				correct = false
				break
			}

			ch2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (ch == ')' && ch2 != '(') ||
				(ch == ']' && ch2 != '[') ||
				(ch == '}' && ch2 != '{') {
				correct = false
				break
			}
		}
	}
	if len(line) > 0 && len(stack) != 0 {
		correct = false
	}

	if correct {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
