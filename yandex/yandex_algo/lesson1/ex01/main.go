package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tr, tc int
	var mode string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &tr, &tc)
	fmt.Fscanln(in, &mode)
	fmt.Fscanln(in, &mode)

	tmp := tr
	switch mode {
	case "freeze":
		if tc < tmp {
			tmp = tc
		}
	case "heat":
		if tc > tmp {
			tmp = tc
		}
	case "auto":
		tmp = tc
	}

	fmt.Println(tmp)
}
