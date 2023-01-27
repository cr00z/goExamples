package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	phones := make([]string, 4)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		fmt.Fscanln(in, &phones[i])
		phones[i] = strings.ReplaceAll(phones[i], "(", "")
		phones[i] = strings.ReplaceAll(phones[i], ")", "")
		phones[i] = strings.ReplaceAll(phones[i], "-", "")
		phones[i] = strings.ReplaceAll(phones[i], "+7", "8")
		if len(phones[i]) == 8 {
			phones[i] = phones[i][1:]
		}
		if len(phones[i]) == 7 {
			phones[i] = "8495" + phones[i]
		}
	}
	for i := 1; i < 4; i++ {
		if phones[i] == phones[0] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
