package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var used []bool
var comps [][]int
var verts [][]int

func dfs(i int) {
	used[i] = true
	lc := len(comps) - 1
	comps[lc] = append(comps[lc], i)
	for _, v := range verts[i] {
		if !used[v] {
			dfs(v)
		}
	}
}

func main() {
	var n, m, v1, v2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	used = make([]bool, n+1)
	comps = make([][]int, 0)
	verts = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		verts[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &v1, &v2)
		if v1 == v2 {
			continue
		}
		verts[v1] = append(verts[v1], v2)
		verts[v2] = append(verts[v2], v1)
	}

	for i := 1; i <= n; i++ {
		if !used[i] {
			comps = append(comps, make([]int, 0))
			dfs(i)
		}
	}

	fmt.Println(len(comps))
	for _, comp := range comps {
		fmt.Println(len(comp))
		fmt.Println(strings.Trim(fmt.Sprint(comp), "[]"))
	}

}
