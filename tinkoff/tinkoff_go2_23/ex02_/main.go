package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	// min
	numOfMax := n % k
	numOfMin := k - numOfMax
	sqOfMin := (n / k) * (n / k)
	sqOfMax := (n/k + 1) * (n/k + 1)
	fmt.Print(numOfMin*sqOfMin+numOfMax*sqOfMax, " ")
	// max
	numOfMax = 1
	numOfMin = k - numOfMax
	sqOfMin = 1
	sqOfMax = (n - numOfMin) * (n - numOfMin)
	fmt.Println(numOfMin*sqOfMin + numOfMax*sqOfMax)
}
