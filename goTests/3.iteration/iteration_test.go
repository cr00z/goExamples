package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	expected := "aaaaaaaaaa"

	get := Repeat("a", 10)

	assert.Equal(t, get, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkRepeatStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStr("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
