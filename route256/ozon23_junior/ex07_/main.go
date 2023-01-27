package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	//now := time.Now()

	var m int
	var n, f1, f2 uint16
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	friends := [50000]map[uint16]struct{}{}
	for i := uint16(0); i < n; i++ {
		friends[i] = make(map[uint16]struct{}, 5)
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &f1, &f2)
		friends[f1-1][f2-1] = struct{}{}
		friends[f2-1][f1-1] = struct{}{}
	}

	for i := uint16(0); i < n; i++ {
		candidates := make(map[uint16]uint16, 25)
		for fr := range friends[i] {
			for frFr := range friends[fr] {
				_, inFriends := friends[frFr][i]
				if frFr != i && !inFriends {
					candidates[frFr]++
				}
			}
		}
		var maxFriends uint16
		for _, val := range candidates {
			if val > maxFriends {
				maxFriends = val
			}
		}
		recomFriends := make([]int, 0, 25)
		for key, val := range candidates {
			if val == maxFriends {
				recomFriends = append(recomFriends, int(key)+1)
			}
		}
		if len(recomFriends) > 0 {
			sort.Ints(recomFriends)
			fmt.Println(strings.Trim(fmt.Sprint(recomFriends), "[]"))
		} else {
			fmt.Println(0)
		}

		//var maxFriends uint16
		//recomFriends := make([]int, 0, 25)
		//for key, val := range candidates {
		//	if val > maxFriends {
		//		maxFriends = val
		//		recomFriends = make([]int, 0, 25)
		//	}
		//	if val >= maxFriends {
		//		recomFriends = append(recomFriends, int(key))
		//	}
		//}
		//
		//if len(recomFriends) > 0 {
		//	sort.Ints(recomFriends)
		//	for _, val := range recomFriends {
		//		fmt.Print(val+1, " ")
		//	}
		//	fmt.Println()
		//} else {
		//	fmt.Println(0)
		//}
	}
	//fmt.Println(time.Since(now))
}
