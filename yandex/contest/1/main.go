package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
func main() {
	var j, s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &j)
	fmt.Fscanln(in, &s)
	jm := make(map[rune]bool)
	count := 0;
	for _, item := range j {
		jm[item] = true
	}
	for _, item := range s {
		if _, inMap := jm[item]; inMap {
			count++
		}
	}
	fmt.Println(count)
}