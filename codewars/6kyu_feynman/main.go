package main

import "fmt"

func CountSquares(n uint) uint {
	var cnt uint = 1
	var i uint = 2
	for ; i <= n; i++ {
		cnt += i * i
	}
	return cnt
}

func main() {
	fmt.Println(CountSquares(8))
}
