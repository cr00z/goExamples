package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
)

type leaf struct {
	firstIdx int
	lastIdx  int
	sum      int
}

func getWeight(cit []int, l *leaf) int {
	penalty := 0
	allSum := 0
	for i := l.firstIdx; i < l.lastIdx; i++ {
		weight := cit[i] - penalty
		if weight < 0 {
			weight = 0
		}
		allSum += weight
		penalty++
	}
	return allSum
}

func main() {
	var n, m, maxSum int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	cit := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &cit[i])
		maxSum += cit[i]
	}

	if maxSum < m {
		fmt.Println(-1)
		os.Exit(0)
	}

	sort.Slice(cit, func(i, j int) bool {
		return cit[i] > cit[j]
	})
	//fmt.Println(cit)

	lst := new(list.List)
	lf := &leaf{
		firstIdx: 0,
		lastIdx:  n,
	}
	lst.PushBack(lf)

	globalSum := getWeight(cit, lf)
	lf.sum = globalSum
	// если за 1 день - сразу выход
	if globalSum >= m {
		fmt.Println(1)
		os.Exit(0)
	}

	// иначе добавляем второй лист
	lf2 := &leaf{
		firstIdx: n,
		lastIdx:  n,
		sum:      0,
	}
	lst.PushBack(lf2)

	for globalSum < m {
		l2Lst := lst.Back()
		lst.Remove(l2Lst)
		l1Lst := lst.Back()
		l2 := l2Lst.Value.(*leaf)
		l1 := l1Lst.Value.(*leaf)
		lst.PushBack(l2)

		oldsum1 := l1.sum
		oldsum2 := l2.sum
		oldsum := oldsum1 + oldsum2
		if l1.lastIdx-l1.firstIdx > 1 {
			l1.lastIdx--
			l2.firstIdx--
		} else {
			if l2.lastIdx-l2.firstIdx == 1 {
				fmt.Println(-1)
				os.Exit(0)
			}
			lst.PushBack(&leaf{
				firstIdx: n,
				lastIdx:  n,
				sum:      0,
			})
			continue
		}
		newsum1 := getWeight(cit, l1)
		newsum2 := getWeight(cit, l2)
		newsum := newsum1 + newsum2

		if oldsum <= newsum {
			l1.sum = newsum1
			l2.sum = newsum2
			globalSum = globalSum - oldsum + newsum
			//fmt.Println(globalSum)
			if globalSum >= m {
				fmt.Println(lst.Len())
				os.Exit(0)
			}
		} else {
			lst.PushBack(&leaf{
				firstIdx: n,
				lastIdx:  n,
				sum:      0,
			})
		}
	}

}
