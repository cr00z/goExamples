package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func scanInts(in io.Reader, n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &result[i])
	}
	return result
}

func checkHearBigger(in []int, out []int) bool {
	for i := 0; i < len(in); i++ {
		if out[i] > in[i] {
			return false
		}
	}
	return true
}

func checkHearEqual(in []int, out []int) bool {
	for i := 0; i < len(in); i++ {
		if out[i] != in[i] {
			return false
		}
	}
	return true
}

func getMaxElementIndex(arr []int) int {
	maxIdx := 0
	maxElem := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxElem {
			maxElem = arr[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func minElem(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var nn, n, nSharps, sharp int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &nn)
LOOP:
	for tst := 0; tst < nn; tst++ {
		fmt.Fscan(in, &n)
		inHears := scanInts(in, n)
		outHears := scanInts(in, n)
		fmt.Fscan(in, &nSharps)
		sharps := make(map[int]int, nSharps)
		for i := 0; i < nSharps; i++ {
			fmt.Fscan(in, &sharp)
			sharps[sharp]++
		}

		// проверка, что длина не больше начальной
		if !checkHearBigger(inHears, outHears) {
			fmt.Println("NO")
			continue
		}
		
		stepQueue := [][2]int{
			{0, n},
		}
		//fmt.Println(inHears, outHears, sharps, stepQueue)

		for len(stepQueue) > 0 {
			step := stepQueue[0]
			stepQueue = stepQueue[1:]
			//fmt.Print(" Step: ", step)

			// проверить что стрижка этого куска закончена
			if checkHearEqual(inHears[step[0]:step[1]], outHears[step[0]:step[1]]) {
				continue
			}

			maxIdx := step[0] + getMaxElementIndex(outHears[step[0]:step[1]])
			hearLen := outHears[maxIdx]
			//fmt.Print(" HearLen: ", hearLen)

			if inHears[maxIdx] != outHears[maxIdx] {
				// нет подходящей бритвы
				if sharps[hearLen] == 0 {
					fmt.Println("NO")
					continue LOOP
				} 
				
				sharps[hearLen]--
				// подстричь
				for i := step[0]; i < step[1]; i++ {
					inHears[i] = minElem(inHears[i], hearLen)
				}
				//fmt.Println(" inHears: ", inHears)				
			}


			start := -1
			stop := -1
			for i := step[0]; i < step[1]; i++ {
				if outHears[i] != hearLen {
					if start == -1 {
						start = i
						stop = i + 1
					} else {
						stop++
					}
				} else {
					if start != stop {
						stepQueue = append(stepQueue, [2]int{start, stop})
						start = -1
						stop = -1
					}
				}
			}
			if start != stop {
				stepQueue = append(stepQueue, [2]int{start, stop})
			}
		}
		//fmt.Println(inHears, outHears, sharps, stepQueue)
		fmt.Println("YES")
	}
}