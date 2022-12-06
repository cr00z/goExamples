package arrays

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3, 4, 5}
		want := 15

		// Act
		got := Sum(numbers)

		// Assert
		assert.Equal(t, got, want)
	})
	t.Run("collection of any size", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3, 4}
		want := 10

		// Act
		got := Sum(numbers)

		// Assert
		assert.Equal(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	// Arrange
	want := []int{3, 9}

	// Act
	got := SumAll([]int{1, 2}, []int{0, 9})

	// Assert
	// assert.Equal(t, got, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9}, []int{4, 5, 6})
	}
}

func BenchmarkSumAll2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll2([]int{1, 2}, []int{0, 9}, []int{4, 5, 6})
	}
}
