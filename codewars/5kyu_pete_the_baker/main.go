package main

import "fmt"

func Cakes(recipe, available map[string]int) int {
	var result int
	first := true
	for ingr, val := range recipe {
		avail_val, ok := available[ingr]
		if ok {
			avail_res := avail_val / val
			if first || (avail_res < result) {
				result = avail_res
				first = false
			}
		} else {
			result = 0
		}
		if result == 0 {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(Cakes(
		map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100},
		map[string]int{"sugar": 500, "flour": 2000, "milk": 2000},
	))
}
