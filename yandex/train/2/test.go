package main

import (
	"fmt"
	"math/rand"
)

func main2() {
	var l1 = []rune(".#")
	var l2 = []string{"left", "right"}
	var l3 = []string{"window", "aisle"}

	fmt.Println(10)
	for i := 0; i < 10; i++ {
		line := make([]rune, 7)
		for i := range line {
			line[i] = l1[rand.Intn(len(l1))]
		}
		line[3] = '_'
		fmt.Println(string(line))
	}
	
	fmt.Println(10)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(3) + 1, l2[rand.Intn(len(l2))], l3[rand.Intn(len(l3))])
	}

}
