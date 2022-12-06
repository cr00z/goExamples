package main

import "fmt"

func remove(slice []interface{}, s int) []interface{} {
	return append(slice[:s], slice[s+1:]...)
}

func Josephus(items []interface{}, k int) []interface{} {
	die := make([]interface{}, 0)
	pos := 1
	for len(items) != 0 {
		if pos == k {
			die = append(die, items[0])
			fmt.Println(die)
			pos = 0
		} else {
			items = append(items, items[0])
		}
		items = remove(items, 0)
		pos++
	}
	return die
}

func main() {
	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7}, 3)...)
}
