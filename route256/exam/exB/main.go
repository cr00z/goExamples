package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValid(day, month, year int) bool {
	mDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	isVis := (year%4 == 0) && (year%100 != 0) || (year%400 == 0)
	if isVis {
		mDays[1] = 29
	}
	if day <= mDays[month-1] {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var day, month, year int
		fmt.Fscan(in, &day, &month, &year)
		if isValid(day, month, year) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
