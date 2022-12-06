package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func timeToInt(time string) (int, bool) {
	numsH := (time[0]-'0')*10 + (time[1] - '0')
	if numsH > 23 {
		return 0, false
	}
	numsM := (time[3]-'0')*10 + (time[4] - '0')
	if numsM > 59 {
		return 0, false
	}
	numsS := (time[6]-'0')*10 + (time[7] - '0')
	if numsS > 59 {
		return 0, false
	}
	return int(numsH)*3600 + int(numsM)*60 + int(numsS), true
}

func main() {
	var packsNum, periodsNum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &packsNum)
	for i := 0; i < packsNum; i++ {
		result := true
		var ticks [86400]byte
		fmt.Fscan(in, &periodsNum)
		j := 0
		var period string
	Loop:
		for ; j < periodsNum; j++ {
			fmt.Fscan(in, &period)
			// fmt.Println(period)
			time := strings.Split(period, "-")
			time0, result1 := timeToInt(time[0])
			if !result1 {
				result = false
				break
			}
			time1, result2 := timeToInt(time[1])
			if !result2 {
				result = false
				break
			}
			result = (time1-time0 >= 0)
			if !result {
				break
			}
			for i := time0; i <= time1; i++ {
				result = (ticks[i] == 0)
				if !result {
					break Loop
				}
				ticks[i] = 1
			}
		}
		if result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		for k := j + 1; k < periodsNum; k++ {
			fmt.Fscan(in, &period)
		}
	}
}
