package wc

import (
    "regexp"
	"strings"
)

// Counter maps words to their counts
type Counter map[string]int

var splitter *regexp.Regexp = regexp.MustCompile(" ")

// WordCountRegexp counts absolute frequencies of words in a string.
// Uses Regexp.Split() to split the string into words.
func WordCountRegexp(s string) Counter {
    counter := make(Counter)
    for _, word := range splitter.Split(s, -1) {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}