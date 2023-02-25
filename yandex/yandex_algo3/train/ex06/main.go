package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type os_struct struct {
	first int
	last  int
}

func main() {
	var m, n, first, last int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &m, &n)
	os := list.New()
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &first, &last)
		if i == 0 {
			os.PushBack(os_struct{first, last})
		} else {
			var next *list.Element
			var insertPos *list.Element
			for e := os.Front(); e != nil; e = next {
				osCurrent := e.Value.(os_struct)
				next = e.Next()

				if osCurrent.first > last && insertPos == nil {
					insertPos = e
					continue
				}

				//fmt.Println(osCurrent.first, osCurrent.last, first, last)

				if (osCurrent.first >= first && osCurrent.first <= last) ||
					(osCurrent.last >= first && osCurrent.last <= last) {
					//fmt.Println("R")
					os.Remove(e)
				}

			}
			if insertPos != nil {
				os.InsertBefore(os_struct{first, last}, insertPos)
			} else {
				os.PushBack(os_struct{first, last})
			}
		}
	}
	fmt.Println(os.Len())
}
