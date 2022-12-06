// echo "rus 2" | go run 1.3.4.go

package main

import (
	"fmt"
)

func main() {
	var text, res string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	if len(text) > width {
		res = text[:width] + "..."
	} else {
		res = text
	}

	fmt.Println(res)
}
