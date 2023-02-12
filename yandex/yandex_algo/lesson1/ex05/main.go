package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var K1, M, K2, P2, N2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &K1, &M, &K2, &P2, &N2)

	fullN2 := (P2-1) * M + P2
	KvOnFloor := K2 / fullN2
	if K2 > KvOnFloor * fullN2 {
		KvOnFloor++
	}

	if (K2 - 1) / KvOnFloor != KvOnFloor * fullN2 {

	}



	numKv2--
	fullNumFloors2 := (pod2-1) * numFloors + floor2
	kvOnFloor := numKv2 / fullNumFloors2
	if numKv2 < kvOnFloor * fullNumFloors2 {
		kvOnFloor++
	}
	if numKv2 / kvOnFloor !=
	p1 := k1/(m*kk) + 1
	fmt.Println(p1)
	n1 := (k1 - (p1-1)*m) / kk
	if k1-(p1-1)*m != n1*kk {
		n1++
	}
	fmt.Println(n1)
}
