package main

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	friends := strings.Split(strings.ToUpper(s), ";")
	friends_map := make(map[string][]string)
	for _, friend := range friends {
		idents := strings.Split(friend, ":")
		friends_map[idents[1]] = append(friends_map[idents[1]], idents[0])
	}

	keys := make([]string, 0, len(friends_map))
	for key := range friends_map {
		sort.Strings(friends_map[key])
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var result string
	for _, surname := range keys {
		for _, name := range friends_map[surname] {
			result += "(" + surname + ", " + name + ")"
		}
	}
	return result
}

func main() {
	s := "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill"
	fmt.Println(Meeting(s))
}
