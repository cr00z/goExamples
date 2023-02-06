package main

import (
	"bufio"
	"fmt"
	"os"
)

func isLetter(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func correct1(idx int, str string) bool {
	if idx+4 > len(str) {
		return false
	}
	if isLetter(str[idx]) && isDigit(str[idx+1]) && isLetter(str[idx+2]) && isLetter(str[idx+3]) {
		return true
	}
	return false
}

func correct2(idx int, str string) bool {
	if idx+5 > len(str) {
		return false
	}
	if isLetter(str[idx]) &&
		isDigit(str[idx+1]) &&
		isDigit(str[idx+2]) &&
		isLetter(str[idx+3]) &&
		isLetter(str[idx+4]) {
		return true
	}
	return false
}

func main() {
	var n int
	var str string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
		//fmt.Println(str)
		idx := 0
		result := ""
		for idx < len(str) {
			if correct1(idx, str) {
				result += str[idx:idx+4] + " "
				idx += 4
				continue
			}
			if correct2(idx, str) {
				result += str[idx:idx+5] + " "
				idx += 5
				continue
			}
			result = "-"
			break
		}
		fmt.Println(result)
	}
}
