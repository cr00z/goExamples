package main

import (
	"fmt"
)

func CreatePhoneNumber(numbers [10]uint) string {
	bytes := []byte{}
	for i := 0; i < 10; i++ {
		bytes = append(bytes, byte(numbers[i])+0x30)
	}
	strnum := string(bytes)
	return "(" + strnum[:3] + ") " + strnum[3:6] + "-" + strnum[6:]
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
