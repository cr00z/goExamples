package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, x, y, minX, minY, maxX, maxY int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &k)
	if k > 0 {
		fmt.Fscan(in, &minX, &minY)
		maxX = minX
		maxY = minY
	}
	for i := 1; i < k; i++ {
		fmt.Fscan(in, &x, &y)
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	fmt.Println(minX, minY, maxX, maxY)
}
