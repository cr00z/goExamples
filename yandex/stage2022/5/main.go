package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Potion struct {
	a int64
	b int64
	c int
}

type GoodPotion struct {
	a int64
	b int64
}

func main() {
	// load
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	// кандидаты в рецепты (некоторые компоненты не расшифрованы)
	items := make(map[int]Potion, n - 2)
	// где используется зелье
	used := make(map[int][]int)
	// готовые рецепты
	goodItems := make(map[int]GoodPotion, n - 2)
	for i := 3; i <= n; i++ {
		p := Potion{}
		var k, item int
		fmt.Fscan(in, &k)
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &item)
			switch item {
			case 1:
				p.a++
			case 2:
				p.b++
			default:
				if val, inMap := goodItems[item]; inMap {
					p.a += val.a
					p.b += val.b
				} else {
					p.c++
					used[item] = append(used[item], i)
				}
			}
			if j == k-1 {
				if p.c > 0 {
					items[i] = p
				} else {
					goodItems[i] = GoodPotion{p.a, p.b}
				}
			}
		}
	}
	// fmt.Println(goodItems)
	// fmt.Println(items)
	// fmt.Println(used)

	// optimize
	keys := make([]int, 0, len(used))
	for key := range used {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	failure := make(map[int]struct{}, n)
	for _, key := range keys {
		// если компонент есть в известных рецептах
		if ab, inMap := goodItems[key]; inMap {
			// смотрим в каких рецептах он используется
			for _, pot := range used[key] {
				p := items[pot]
				p.a += ab.a
				p.b += ab.b
				// добавляем его состав в нерасшифрованные рецепты
				items[pot] = p
			}
		} else {
			// если компонента нет в известных - добавляем все рецепты в неизвестные
			for _, pot := range used[key] {
				failure[pot] = struct{}{}
			}	
		}	
	}

	// все расшифрованные рецепты добавляем в известные
	for key, val := range items {
		if _, inMap := failure[key]; !inMap {
			goodItems[key] = GoodPotion{val.a, val.b}
		}
	}

	// queries
	var m, qp int
	var qa, qb int64
	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &qa, &qb, &qp)
		if val, inMap := goodItems[qp]; inMap {
			if val.a <= qa && val.b <= qb {
				fmt.Print(1)
				continue
			}
		}
		fmt.Print(0)
	}
}