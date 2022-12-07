package main

import (
	"bufio"
	"fmt"
	"os"
)

type Leaf struct {
	num int
	p *Leaf
	vl *Leaf
	vr *Leaf
}

func printTree(root *Leaf) {
	if root.vl != nil {
		printTree(root.vl)
	}
	fmt.Print(root.num, " ")
	if root.vr != nil {
		printTree(root.vr)		
	}
}

func main() {
	var n, q int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &q)
	tree := make(map[int]*Leaf, n)
	for i := 1; i <= n; i++ {
		var p *Leaf
		if i != 1 {
			p = tree[i/2]		
		}
		tree[i] = &Leaf{
			num: i,
			p: p,
		}
		if i != 1 {
			if i % 2 == 0 {
				tree[i/2].vl = tree[i]
			} else {
				tree[i/2].vr = tree[i]
			}				
		}
	}
	root := 1

	var num int
	for i := 0; i < 1; i++ {
		fmt.Fscan(in, &num)

		if num == root {
			continue
		}
		pNum := tree[num].p.num
		var ppNum int
		if pNum != root {
			ppNum = tree[pNum].p.num
			if pNum == tree[ppNum].vl.num {
				tree[ppNum].vl = tree[num]
			} else {
				tree[ppNum].vr = tree[num]
			}
			tree[num].p = tree[ppNum]
		}
		tree[pNum].p = tree[num]
		if num == tree[pNum].vl.num {
			tree[num].vl, tree[pNum].vl = tree[pNum].vl, tree[num].vl
		} else {
			tree[num].vr, tree[pNum].vr = tree[pNum].vr, tree[num].vr
		}
		if pNum != root {
			tree[num].p = tree[ppNum]
		} else {
			root = num
		}

		fmt.Println(num, pNum, ppNum)
	}

	//printTree(tree[root])
}