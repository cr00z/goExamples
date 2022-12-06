package main

import "fmt"

func PageDigits(pages uint64) uint64 {
	var summary uint64 = 0
	var i uint64 = 1
	for pages != 0 {
		summary += pages
		if pages > 9*i {
			pages = pages - 9*i
		} else {
			pages = 0
		}
		i *= 10
	}
	return summary
}

func main() {
	fmt.Println(PageDigits(99999))
}
