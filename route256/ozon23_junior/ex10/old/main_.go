package old

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

type pos struct {
	first, last int
}

func main_() {
	var n, m int
	var str string

	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	vocab := make([]string, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
		vocab[i] = reverse(str)
	}
	sort.Strings(vocab)

	keys := make(map[byte]pos)
	prev := byte(0)
	for idx, str := range vocab {
		if str[0] != prev {
			prevPos := keys[prev]
			prevPos.last = idx
			keys[prev] = prevPos

			currPos := keys[str[0]]
			currPos.first = idx
			keys[str[0]] = currPos

			prev = str[0]
		}
	}

	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &str)
		str = reverse(str)

		idx, ok := keys[str[0]]
		if !ok {
			fmt.Println(reverse(vocab[0]))
			continue
		}

		if idx.last == 0 {
			idx.last = len(vocab)
		}

		//fmt.Println(idx)
		k := 0
		candIdx := idx.first

		for p := idx.first; p < idx.last; p++ {
			if str == vocab[p] {
				continue
			}
			if str[k] == vocab[p][k] {
				k++
				candIdx = p
				continue
			}
			if str[k] > vocab[p][k] {
				break
			}
		}

		if str == vocab[candIdx] {
			candIdx++
			if candIdx == len(vocab) {
				candIdx = 0
			}
		}

		fmt.Println(reverse(vocab[candIdx]))
	}
	//fmt.Println(vocab)
}
