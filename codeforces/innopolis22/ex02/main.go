package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

// func int2min(num int64, lim int64) (int64, bool) {
// 	if num < 10 && num > lim {
// 		return num, true
// 	}
// 	var result int64 = num
// 	digitsNum := []byte(strconv.FormatInt(num, 10))
// 	lenNum := len(digitsNum)
// 	for i := 0; i < lenNum; i++ {
// 		for j := i+1; j < lenNum; j++ {
// 			// if digitsNum[i] >= digitsNum[j] {
// 			// 	continue
// 			// }
// 			digitsNum[i], digitsNum[j] = digitsNum[j], digitsNum[i]
// 			newNum, _ := strconv.ParseInt(string(digitsNum), 10, 64)
// 			if newNum < result && newNum > lim {
// 				result = newNum
// 			}
// 			digitsNum[i], digitsNum[j] = digitsNum[j], digitsNum[i]
// 		}
// 	}
// 	if result > lim {
// 		return result, true
// 	}
// 	return 0, false
// }

func int2min(num []byte, lim int64) (int64, bool) {
	// if num < 10 && num > lim {
	// 	return num, true
	// }
	var result int64
	sort.Slice(num, func(i int, j int)bool {return num[i] < num[j]})
	result, _ = strconv.ParseInt(string(num), 10, 64)
	lenNum := len(num)
	for i := 0; i < lenNum; i++ {
		for j := i+1; j < lenNum; j++ {
			// if digitsNum[i] >= digitsNum[j] {
			// 	continue
			// }
			num[i], num[j] = num[j], num[i]
			newNum, _ := strconv.ParseInt(string(num), 10, 64)
			fmt.Println("Conv:", string(num), newNum)
			if newNum < result && newNum > lim {
				result = newNum
			}
			num[i], num[j] = num[j], num[i]
		}
	}
	if result > lim {
		return result, true
	}
	return 0, false
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	emp := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &emp[i])
	}

	fmt.Println(emp)
	
	maxLimit := 1
	for i := 0; (i < n) && (n - i > maxLimit); i++ {
		prevNums := []byte(emp[i])
		sort.Slice(prevNums, func(i int, j int)bool {return prevNums[i] < prevNums[j]})
		prev, _ := strconv.ParseInt(string(prevNums), 10, 64)
		fmt.Printf("Prev: %d\n", prev)

		limit := 1
		for j := i+1; j < n; j++ {
			next, isOk := int2min([]byte(emp[j]), prev)
			fmt.Println(emp[j], prev, next, isOk)
			if isOk {
				limit++
				prev = next
			}
		}
		if limit > maxLimit {
			maxLimit = limit
		}
	}
	fmt.Println(maxLimit)
}