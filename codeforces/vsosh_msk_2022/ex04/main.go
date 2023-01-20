package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var l, r, ls, rs, maxLeft, minRight int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &l, &r)
	leftM := make(map[int]struct{}, l)
	rightM := make(map[int]struct{}, r)
	for i := 0; i < l; i++ {
		fmt.Fscan(in, &ls)
		leftM[ls] = struct{}{}
		if i == 0 {
			maxLeft = ls
		} else {
			if ls > maxLeft {
				maxLeft = ls
			}
		}
	}
	for i := 0; i < r; i++ {
		fmt.Fscan(in, &rs)
		rightM[rs] = struct{}{}
		if i == 0 {
			minRight = ls
		} else {
			if ls < minRight {
				minRight = ls
			}
		}
	}

	leftOdd := make([]int, 0, len(leftM))
	leftEven := make([]int, 0, len(leftM))
	rightOdd := make([]int, 0, len(rightM))
	rightEven := make([]int, 0, len(rightM))
	torpedos := make([]int, 0)
	for key := range leftM {
		if key < minRight {
			torpedos = append(torpedos, 0, key)
		} else {
			if key % 2 == 0 {
				leftOdd = append(leftOdd, key)
			} else {
				leftEven = append(leftEven, key)
			}
		}
	}
	for key := range rightM {
		if key > maxLeft {
			torpedos = append(torpedos, 0, key)
		} else {
			if key % 2 == 0 {
				rightOdd = append(rightOdd, key)
			} else {
				rightEven = append(rightEven, key)
			}
		}
	}
	sort.Ints(leftOdd)
	sort.Ints(leftEven)
	sort.Ints(rightOdd)
	sort.Ints(rightEven)
	
	posl := 0
	posr := 0
	

	for posl < len(left) && posr < len(right) && left[posl] >= right[posr] {
		diff := left[posl] - right[posr]
		if diff % 2 == 0 {
			torpedos = append(torpedos, diff / 2, left[posl] + right[posr] / 2)
		}
		
	}
}