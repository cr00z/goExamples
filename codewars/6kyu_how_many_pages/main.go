package main

import "fmt"

func AmountOfPages(summary int) int {
	var result int
	if summary <= 9 {
		result = summary
	} else if summary <= 189 {
		result = 9 + (summary-9)/2
	} else if summary <= 2889 {
		result = 99 + (summary-189)/3
	} else if summary <= 38889 {
		result = 999 + (summary-2889)/4
	} else if summary <= 488889 {
		result = 9999 + (summary-38889)/5
	} else {
		result = 100000
	}
	return result
}

func main() {
	fmt.Println(AmountOfPages(1095))
}
