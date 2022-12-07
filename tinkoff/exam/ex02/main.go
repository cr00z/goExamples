package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, x, y, z, i, j, k, kol int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &c, &x, &y, &z)
	sum := a * x + b * y + c * z
	for i = 0; i <= sum / a; i++ {
		if i * a > sum {
			break
		}
		for j = 0; j <= sum / b; j++ {
			if i * a + j * b > sum {
				break
			}
			for k = 0; k <= sum / c; k++ {
				sumk := i * a + j * b + k * c
				if sumk > sum {
					break
				}
				if sumk == sum {
					//fmt.Println(i, j, k)
					kol++
				}
			}
		}
	}
	fmt.Println(kol)
}