package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num, time, out, steps int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &num, &time)
	emp := make([]int, num, num)
	for i := 0; i < num; i++ {
		fmt.Fscan(in, &emp[i])
	}
	fmt.Fscan(in, &out)

	steps = emp[num-1] - emp[0]
	up := emp[out-1] - emp[0]
	down := emp[num-1] - emp[out-1]
	if (up > time) && (down > time) {
		if up > down {
			steps += down
		} else {
			steps += up
		}
	}
	fmt.Println(steps)
}
