package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

func toDivs(n int64) ([]int, error) {
	divisors := []int64{9, 8, 7, 6, 5, 3, 2}
	result := make([]int, 0)
	for _, div := range divisors {
		for n % div == 0 {
			result = append(result, int(div))
			n = n / div
		}
	}
	if n > 1 {
		return nil, errors.New("bad number")
	}
	return result, nil
}


func main() {
	var n int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	nums, err := toDivs(n)
	if err != nil {
		fmt.Println(-1)
	} else {
		sort.Ints(nums)
		for _, num := range nums {
			fmt.Print(num)
		}
		
	}
}