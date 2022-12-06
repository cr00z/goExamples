package main

import "fmt"

func Zeros(n int) int {
	res := 0
	for n > 0 {
		n /= 5
		res += n
	}
	return res
}

func main() {
	fmt.Println(Zeros(30))
}
