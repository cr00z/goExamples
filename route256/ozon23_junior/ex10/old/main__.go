package old

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func cmp(bytes1 []byte, bytes2 []byte) bool {

	for i := 0; i < 10; i++ {
		if i == len(bytes1) {
			return false
		}
		if bytes1[i] != bytes2[i] {
			return false
		}
	}
	return true
}

type node map[byte]struct {
	Node node
	Stop bool
}

func main__() {
	var n, m int
	var str string

	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	root := make(map[byte]struct {
		Node node
		Stop bool
	}, 26)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
		str = reverse(str)
		current := root
		for i := 0; i < len(str); i++ {
			key := str[i]
			if next, ok := current[key]; ok {
				current = next.Node
			} else {
				tmp := struct {
					Node node
					Stop bool
				}{}
				tmp.Node = make(map[byte]struct {
					Node node
					Stop bool
				}, 26)
				if i == len(str)-1 {
					tmp.Stop = true
				}
				current[key] = tmp
				current = current[key].Node
			}
		}
	}

	var levels [10]map[byte]struct {
		Node node
		Stop bool
	}
	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &str)
		str = reverse(str)
		current := root
		rifma := make([]byte, 0, 10)

		i := 0
		for ; i < len(str); i++ {
			levels[i] = current
			key := str[i]
			if next, ok := current[key]; ok {
				rifma = append(rifma, key)
				current = next.Node
				continue
			}
			break
		}
		for {
			if len(current) > 0 {
				for key, next := range current {
					rifma = append(rifma, key)
					current = next.Node
					i++
					break
				}
				continue
			}
			break
		}

		i--
		//fmt.Println(i)
		if str == string(rifma) {
			if len(levels[i]) < 2 {
				rifma = rifma[:i]
				i--
			}
			var stop bool
			for ; len(levels[i]) < 2; i-- {
				if levels[i][rifma[i]].Stop {
					stop = true
					break
				}
				rifma = rifma[:i]
			}
			if !stop {
				for key, next := range levels[i] {
					if key != rifma[i] {
						rifma[i] = key
						current = next.Node
						break
					}
				}
				for len(current) > 0 {
					for key, next := range current {
						rifma = append(rifma, key)
						current = next.Node
						break
					}
				}
			}
		}

		fmt.Println(reverse(string(rifma)))
	}
}
