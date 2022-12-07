package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, k, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)

	s := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		pow := 1
		for a != 0 {
			key := a % 10
			if key < 9 {
				key = (9 - key) * pow
				s[key] = s[key] + 1
			}
			pow *= 10
			a /= 10
		}
	}

	keys := make([]int, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int)bool {
		return keys[i] > keys[j]
	})

	sum := 0
	for _, key := range keys {
		val := s[key]
		if val > k {
			val = k
		}
		sum += key * val
		k = k - val
		if k == 0 {
			break
		}
	}
	fmt.Println(sum)
}