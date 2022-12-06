package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, rNum, req int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &rNum)
		//fmt.Println(fNum)
		var first, second int
		var len int
		if rNum == 1 {
			fmt.Fscan(in, &first)
			len = 1
		} else if rNum == 2 {
			fmt.Fscan(in, &first)
			fmt.Fscan(in, &second)
			len = 2
		} else {
			fmt.Fscan(in, &first)
			fmt.Fscan(in, &second)
			firstN := 1
			secondN := 1
			len = 2
			for first == second && rNum > 2 {
				fmt.Fscan(in, &second)
				firstN++
				secondN = 1
				rNum--
			}
			prev := second
			for j := 0; j < rNum-2; j++ {
				//fmt.Println(first, second, firstN, secondN)
				fmt.Fscan(in, &req)
				if req == first {
					firstN++
				} else if req == second {
					secondN++
				} else {
					if firstN+secondN > len {
						len = firstN + secondN
					}
					first = prev
					second = req
					firstN = 1
					secondN = 1
				}
				prev = req
			}
			if firstN+secondN > len {
				len = firstN + secondN
			}
		}
		fmt.Println(len)
	}
}
