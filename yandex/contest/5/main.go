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
	var a, b string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &a)
	fmt.Fscanln(in, &b)
	if SortString(a) == SortString(b) {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}