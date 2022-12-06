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

func validate(periods *[]string) bool {
	times := make([][2]int, 0, len(*periods))
	for _, period := range *periods {
		time := strings.Split(period, "-")
		time0, correct := timeToInt(time[0])
		if !correct {
			return false
		}
		time1, correct := timeToInt(time[1])
		if !correct {
			return false
		}
		duration := time1 - time0
		if duration < 0 {
			return false
		}
		leftAppend := false
		rightAppend := false
		var pos int
		for idx, inter := range times {
			if !((inter[0] < time0 && inter[1] < time0) || (inter[0] > time1 && inter[1] > time1)) {
				return false
			}
			if time0 == inter[1]+1 {
				leftAppend = false
				rightAppend = true
				pos = idx
			}
			if time1 == inter[0]-1 {
				rightAppend = false
				leftAppend = true
				pos = idx
			}
		}
		if leftAppend || rightAppend {
			if leftAppend {
				times[pos][0] = time0
			}
			if rightAppend {
				times[pos][1] = time1
			}
		} else {
			times = append(times, [2]int{time0, time1})
		}
	}
	//fmt.Println(times)
	return true
}

func main() {
	var packsNum, periodsNum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &packsNum)
	for i := 0; i < packsNum; i++ {
		fmt.Fscan(in, &periodsNum)
		periods := make([]string, periodsNum)
		for j := 0; j < periodsNum; j++ {
			fmt.Fscan(in, &periods[j])
		}
		if validate(&periods) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
