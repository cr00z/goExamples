package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type hotel struct {
	Num   int
	Mark  int
	Stars int
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var n2 int
		fmt.Fscan(in, &n2)
		var hotels []hotel
		for j := 0; j < n2; j++ {
			var mark int
			fmt.Fscan(in, &mark)
			hotels = append(hotels, hotel{j, mark, -1})
		}
		good := false
		if n2 >= 15 {
			good = true
			sort.Slice(hotels, func(i, j int) bool {
				return hotels[i].Mark > hotels[j].Mark
			})
			max_cnt := 1
			cnt := 0
			stars := 5
			hotels[0].Stars = stars
			mark := hotels[0].Mark
			for i := 1; i < n2; i++ {
				if stars == 1 {
					max_cnt--
				}
				if max_cnt > n2-i {
					good = false
				}
				if cnt == 0 && hotels[i].Mark == mark {
					hotels[i].Stars = stars
					max_cnt++
					continue
				}
				if cnt == 0 && stars > 1 {
					max_cnt++
					cnt = max_cnt
					stars--
				}
				hotels[i].Stars = stars
				cnt--
				mark = hotels[i].Mark
			}

		}
		sort.Slice(hotels, func(i, j int) bool {
			return hotels[i].Num < hotels[j].Num
		})
		for _, hotel := range hotels {
			if good {
				fmt.Printf("%d ", hotel.Stars)
			} else {
				fmt.Print("-1 ")
			}
		}
		fmt.Println()
	}
}
