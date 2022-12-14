package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	expected := 4

	sum := Add(2, 2)

	assert.Equal(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
