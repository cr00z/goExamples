package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type inter struct {
	left  time.Time
	right time.Time
}

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)

	for ti := 0; ti < t; ti++ {
		fmt.Fscan(in, &n)

		valid := "YES"
		inters := make([]inter, n, n)

		for i := 0; i < n; i++ {
			var intStr string
			fmt.Fscan(in, &intStr)
			dates := strings.Split(intStr, "-")

			time1, err1 := time.Parse("15:04:05", dates[0])
			time2, err2 := time.Parse("15:04:05", dates[1])
			if err1 != nil || err2 != nil || time1.After(time2) {
				valid = "NO"
			}

			inters[i] = inter{
				left:  time1,
				right: time2,
			}
		}

		if (n > 1) && (valid == "YES") {
			sort.Slice(inters, func(i, j int) bool {
				return inters[i].left.Before(inters[j].left)
			})

			for j := 0; j < n-1; j++ {
				if !inters[j].right.Before(inters[j+1].left) {
					valid = "NO"
					break
				}
			}
		}

		fmt.Println(valid)
	}
}
