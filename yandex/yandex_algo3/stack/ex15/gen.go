package main

import (
	"os"
	"strconv"
)

func main() {
	file, _ := os.Create("testbig")
	file.WriteString(strconv.Itoa(100_000) + "\n")
	for i := 1; i <= 100_000; i++ {
		file.WriteString(strconv.Itoa(i) + " ")
	}
	file.Close()
}
