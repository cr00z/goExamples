package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	first := make([]int, 5)
	second := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscan(in, &first[i])
	}
	for i := 0; i < 5; i++ {
		fmt.Fscan(in, &second[i])
	}

	botva := true
	for i := 0; i < 1_000_000; i++ {
		cardFirst := first[0]
		first = first[1:len(first)]
		cardSecond := second[0]
		second = second[1:len(second)]

		if (cardFirst == 0 && cardSecond == 9) ||
			(cardFirst > cardSecond && !(cardFirst == 9 && cardSecond == 0)) {
			first = append(first, cardFirst, cardSecond)
		} else {
			second = append(second, cardFirst, cardSecond)
		}

		if len(first) == 0 {
			fmt.Println("second", i+1)
			botva = false
			break
		}
		if len(second) == 0 {
			fmt.Println("first", i+1)
			botva = false
			break
		}
	}

	if botva {
		fmt.Println("botva")
	}
}
