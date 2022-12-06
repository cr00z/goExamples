package main

import (
	"testing"
)

func _Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func _TestSumZero(t *testing.T) {
	if Sum() != 0 {
		t.Errorf("Expected Sum() == 0")
	}
}

func _TestSumOne(t *testing.T) {
	if Sum(1) != 1 {
		t.Errorf("Expected Sum(1) == 1")
	}
}

func _TestSumPair(t * testing.T) {
	if Sum(1, 2) != 3 {
		t.Errorf("Expected Sum(1, 2) == 3")
	}
}

func _TestSumMany(t *testing.T) {
	if Sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}