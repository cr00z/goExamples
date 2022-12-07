package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	sm := float64(n) / 4 / math.Tan(math.Pi / float64(n))
	//central := sm / float64(n) * 3
	fmt.Println(sm)
}