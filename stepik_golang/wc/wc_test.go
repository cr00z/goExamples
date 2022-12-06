// -- wc_test.go --
package wc

import (
    "fmt"
    "math/rand"
    "strings"
    "testing"
)

func BenchmarkRegexp(b *testing.B) {
    for _, length := range []int{10, 100, 1000, 10000} {
        rand.Seed(0)
        phrase := randomPhrase(length)
        name := fmt.Sprintf("Regexp-%d", length)
        b.Run(name, func(b *testing.B) {
            for n := 0; n < b.N; n++ {
                WordCountRegexp(phrase)
            }
        })
    }
}

// randomPhrase returns a phrase of n random words
func randomPhrase(n int) string {
    words := make([]string, n)
	for i := 0; i < n; i++ {
		col := rand.Intn(10)
		bstr := make([]byte, col)
		for i := 0; i<col; i++ {
			bstr[i] = byte(rand.Intn(26)) + 'a'
		}
		words[i] = string(bstr)
	}
	return strings.Join(words, " ")
}