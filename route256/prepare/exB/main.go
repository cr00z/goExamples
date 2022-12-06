package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var packsNum, elemNum, elem int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &packsNum)
	for i := 0; i < packsNum; i++ {
		fmt.Fscan(in, &elemNum)
		foods := make(map[int]int)
		for j := 0; j < elemNum; j++ {
			fmt.Fscan(in, &elem)
			if _, inMap := foods[elem]; inMap {
				foods[elem]++
			} else {
				foods[elem] = 1
			}
		}
		sum := 0
		for price, val := range foods {
			sum += ((val/3)*2 + val%3) * price
		}
		fmt.Println(sum)
	}
}
