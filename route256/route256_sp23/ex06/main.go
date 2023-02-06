package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t, k int
	var printed string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for tt := 0; tt < t; tt++ {
		fmt.Fscan(in, &k)
		fmt.Fscan(in, &printed)
		prMap := make(map[int]struct{}, k)
		diaps := strings.Split(printed, ",")
		for _, diap := range diaps {
			idxs := strings.Split(diap, "-")
			if len(idxs) == 1 {
				idx, _ := strconv.Atoi(idxs[0])
				prMap[idx] = struct{}{}
			} else {
				idx1, _ := strconv.Atoi(idxs[0])
				idx2, _ := strconv.Atoi(idxs[1])
				for idx := idx1; idx <= idx2; idx++ {
					prMap[idx] = struct{}{}
				}
			}
		}

		//fmt.Println(prMap)
		first := -1
		last := -1
		var nfp bool
		for i := 1; i <= k; i++ {
			if _, ok := prMap[i]; !ok {
				if first == -1 {
					first = i
					last = i
				}
				last = i
			} else {
				if first != -1 {
					if nfp {
						fmt.Print(",")
					}
					if first == last {
						fmt.Print(first)
						nfp = true
					} else if first+1 == last {
						fmt.Print(first, ",", last)
						nfp = true
					} else {
						fmt.Print(first, "-", last)
						nfp = true
					}
					first = -1
					last = -1
				}
			}
		}
		if first != -1 {
			if nfp {
				fmt.Print(",")
			}
			if first == last {
				fmt.Print(first)
			} else if first+1 == last {
				fmt.Print(first, ",", last)
			} else {
				fmt.Print(first, "-", last)
			}
		}
		fmt.Println()
	}

}
