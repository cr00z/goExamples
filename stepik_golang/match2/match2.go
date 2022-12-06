package main

import (
    "strings"
	"fmt"
	"regexp"
)

// MatchContains returns true if the string
// contains the pattern, false otherwise.
func MatchContains(pattern string, src string) bool {
    return strings.Contains(src, pattern)
}


// MatchRegexp returns true if the string
// matches the regexp pattern, false otherwise.
func MatchRegexp(pattern string, src string) bool {
    re, err := regexp.Compile(pattern)
    if err != nil {
        return false
    }
    return re.MatchString(src)
}


func main() {
	s := "go is awesome"

	fmt.Println(MatchContains("is", s))
	// true
	fmt.Println(MatchContains("go.*awesome", s))
	// false

	fmt.Println(MatchRegexp("go.*awesome", s))
	// true
	fmt.Println(MatchRegexp("^go", s))
	// true
	fmt.Println(MatchRegexp("awesome$", s))
	// true
}