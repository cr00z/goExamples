package main

import "fmt"

type IntSet struct {
	elems map[int]int
}

func MakeIntSet() IntSet {
	elems := make(map[int]int)
	return IntSet{elems}
}

func (s IntSet) Contains(val int) bool {
	if _, isOk := s.elems[val]; isOk {
		return true
	}
	return false
}

func (s IntSet) Add(val int) bool {
	if s.Contains(val) {
		return false
	}
	s.elems[val] = 0
	return true
}

func main() {
	set := MakeIntSet()

	set.Add(5)
	fmt.Println(set.Contains(5))
	// true

	fmt.Println(set.Contains(42))
	// false

	// элементы множества уникальны,
	// добавить дважды один и тот же элемент не получится
	added := set.Add(5)
	fmt.Println(added)
	// false
}