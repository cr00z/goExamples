package main

import (
	"bufio"
	"fmt"
	"os"
)

type town struct {
	boxes int64
	down  []*town
}

func drop(town *town, k int, x int) {
	town.boxes += int64(x)
	if k > 0 {
		for _, down := range town.down {
			if down != nil {
				drop(down, k-1, x)
			}
		}
	}
}

func main() {
	var n, t1, t2, q, v, k, x int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	towns := make([]*town, n+1)
	for i := 1; i <= n; i++ {
		t := town{
			down: make([]*town, 0),
		}
		towns[i] = &t
	}
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &t1, &t2)
		//fmt.Println(t1, t2)
		towns[t1].down = append(towns[t1].down, towns[t2])
	}
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &v, &k, &x)
		//fmt.Println(v, k, x)
		drop(towns[v], k, x)
	}
	for i := 1; i <= n; i++ {
		fmt.Print(towns[i].boxes, " ")
	}
}
