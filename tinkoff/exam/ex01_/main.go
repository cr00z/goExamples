package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var str2 string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	fmt.Fscanln(in)
	str1, _, _ := in.ReadLine()
	fmt.Fscanln(in, &str2)
	var pos, fails int
	for _, str := range strings.Split(string(str1), " ") {
		colors := str2[pos:pos+len(str)]
		if strings.Contains(colors, "BB") || strings.Contains(colors, "YY") {
			fails++
		}
		pos += len(str)
	}
	fmt.Println(fails)
}