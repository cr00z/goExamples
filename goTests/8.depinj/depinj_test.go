package depinj

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	want := "Hello, Chris"

	Greet(&buffer, "Chris")
	got := buffer.String()

	assert.Equal(t, got, want)
}
