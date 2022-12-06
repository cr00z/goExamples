// echo "a 5" | go run 1.2.8.go

package main

import (
	"fmt"
)

func main() {
	var source, result string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
    // ...
	for times > 0 {
		result += source
		times--
	}

	fmt.Println(result)
}